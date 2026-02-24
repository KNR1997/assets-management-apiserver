package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knr1997/assets-management-apiserver/internal/store"
)

type supplierKey string

const supplierCtx supplierKey = "supplier"

func getsupplierFromCtx(r *http.Request) *store.Supplier {
	supplier, _ := r.Context().Value(supplierCtx).(*store.Supplier)
	return supplier
}

func (app *application) supplierContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "supplierID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		supplier, err := app.store.Supplier.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, supplierCtx, supplier)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type CreatesupplierPayload struct {
	Name string `json:"name" validate:"required,max=100"`
	// Notes string `json:"description"`
}

func (app *application) createSupplierHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatesupplierPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Supplier{
		Name: payload.Name,
	}

	ctx := r.Context()

	if err := app.store.Supplier.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdatesupplierPayload struct {
	Name  *string `json:"name" validate:"omitempty,max=100"`
	Notes *string `json:"notes"`
}

func (app *application) updateSupplier(ctx context.Context, supplier *store.Supplier) error {
	if err := app.store.Supplier.Update(ctx, supplier); err != nil {
		return err
	}

	return nil
}

func (app *application) updateSupplierHandler(w http.ResponseWriter, r *http.Request) {
	supplier := getsupplierFromCtx(r)

	var payload UpdatesupplierPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name != nil {
		supplier.Name = *payload.Name
	}

	ctx := r.Context()

	if err := app.updateSupplier(ctx, supplier); err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, supplier); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getSupplierHandler(w http.ResponseWriter, r *http.Request) {
	supplier := getsupplierFromCtx(r)

	response := TosupplierResponse(supplier)

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type supplierResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

func TosupplierResponse(a *store.Supplier) supplierResponse {
	return supplierResponse{
		ID:   a.ID,
		Name: a.Name,
	}
}

func TosupplierResponseList(categories []store.Supplier) []supplierResponse {
	responses := make([]supplierResponse, 0, len(categories))

	for i := range categories {
		responses = append(responses, TosupplierResponse(&categories[i]))
	}

	return responses
}

func (app *application) getAllSupplierHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := app.store.Supplier.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := TosupplierResponseList(categories)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *application) deleteSupplierHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "supplierID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Supplier.Delete(ctx, id); err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			app.notFoundResponse(w, r, err)
		default:
			app.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
