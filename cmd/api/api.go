package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/knr1997/assets-management-apiserver/internal/auth"
	"github.com/knr1997/assets-management-apiserver/internal/store"
	httpSwagger "github.com/swaggo/http-swagger"
	"go.uber.org/zap"
)

type application struct {
	config        config
	store         store.Storage
	logger        *zap.SugaredLogger
	authenticator auth.Authenticator
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	frontendURL string
	auth        authConfig
}

type authConfig struct {
	basic basicConfig
	token tokenConfig
}

type basicConfig struct {
	user string
	pass string
}

type tokenConfig struct {
	secret string
	exp    time.Duration
	iss    string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *application) mount() http.Handler {
	r := chi.NewRouter()

	// CORS configuration
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Maximum value not ignored by browsers
	}))

	// A good base middleware stack
	r.Use(middleware.RequestID) // important for rate limiting
	r.Use(middleware.RealIP)    // import for rate limiting and analytics and tracing
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer) // recover from crashes
	r.Use(middleware.Timeout(60 * time.Second))

	// Swagger UI
	r.Get("/swagger/*", httpSwagger.WrapHandler)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("all good"))
	})

	r.Route("/api/categories", func(r chi.Router) {
		r.Use(app.AuthTokenMiddleware)
		r.Use(app.AuditContextMiddleware)

		r.Get("/", app.getAllCategoryHandler)
		r.Post("/", app.createCategoryHandler)

		r.Route("/{categoryID}", func(r chi.Router) {
			r.Use(app.categoryContextMiddleware)
			r.Get("/", app.getCategoryHandler)

			r.Patch("/", app.updateCategoryHandler)
			r.Delete("/", app.deleteCategoryHandler)
		})
	})

	r.Route("/api/departments", func(r chi.Router) {
		r.Get("/", app.getAlldepartmentHandler)
		r.Post("/", app.createdepartmentHandler)

		r.Route("/{departmentID}", func(r chi.Router) {
			r.Use(app.departmentContextMiddleware)
			r.Get("/", app.getdepartmentHandler)

			r.Patch("/", app.updatedepartmentHandler)
			r.Delete("/", app.deletedepartmentHandler)
		})
	})

	r.Route("/api/suppliers", func(r chi.Router) {
		r.Get("/", app.getAllSupplierHandler)
		r.Post("/", app.createSupplierHandler)

		r.Route("/{supplierID}", func(r chi.Router) {
			r.Use(app.supplierContextMiddleware)
			r.Get("/", app.getSupplierHandler)

			r.Patch("/", app.updateSupplierHandler)
			r.Delete("/", app.deleteSupplierHandler)
		})
	})

	r.Route("/api/models", func(r chi.Router) {
		r.Get("/", app.getAllModelHandler)
		r.Post("/", app.createModelHandler)

		r.Route("/{modelID}", func(r chi.Router) {
			r.Use(app.modelContextMiddleware)
			r.Get("/", app.getModelHandler)

			r.Patch("/", app.updateModelHandler)
			r.Delete("/", app.deleteModelHandler)
		})
	})

	r.Route("/api/manufacturers", func(r chi.Router) {
		r.Get("/", app.getAllManufacturerHandler)
		r.Post("/", app.createManufacturerHandler)

		r.Route("/{manufacturerID}", func(r chi.Router) {
			r.Use(app.manufacturerContextMiddleware)
			r.Get("/", app.getManufacturerHandler)

			r.Patch("/", app.updateManufacturerHandler)
			r.Delete("/", app.deleteManufacturerHandler)
		})
	})

	r.Route("/api/assets", func(r chi.Router) {
		r.Get("/", app.getAllAssetHandler)
		r.Post("/", app.createAssetHandler)

		r.Route("/{assetID}", func(r chi.Router) {
			r.Use(app.assetContextMiddleware)
			r.Get("/", app.getAssetHandler)

			r.Patch("/", app.updateAssetHandler)
			r.Delete("/", app.deleteAssetHandler)

			r.Post("/checkout", app.checkoutAssetHandler)
			r.Post("/checkin", app.checkinAssetHandler)
		})
	})

	r.Route("/asset-assignments", func(r chi.Router) {
		r.Post("/", app.CreateAssetAssignmentHandler)
	})

	r.Route("/api/profile", func(r chi.Router) {
		r.Use(app.AuthTokenMiddleware)
		r.Patch("/", app.updateUserHandler)
	})

	r.Route("/api/me", func(r chi.Router) {
		r.Use(app.AuthTokenMiddleware)
		r.Get("/", app.meDetailsHandler)
	})

	r.Route("/api/users", func(r chi.Router) {
		r.Use(app.AuthTokenMiddleware)
		r.Get("/", app.getAllUserHandler)
		// r.Post("/", app.createAssetHandler)

		r.Route("/{userID}", func(r chi.Router) {
			r.Use(app.userContextMiddleware)
			// r.Get("/", app.getAssetHandler)

			r.Patch("/", app.updateUserHandler)
			// r.Delete("/", app.deleteAssetHandler)
		})
	})

	// Public routes
	r.Route("/api/authentication", func(r chi.Router) {
		r.Post("/user", app.registerUserHandler)
		r.Post("/token", app.createTokenHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {
	srv := http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	app.logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	app.logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
