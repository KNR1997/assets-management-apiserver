package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/knr1997/assets-management-apiserver/internal/api/requests"
	"github.com/knr1997/assets-management-apiserver/internal/api/responses"
	"github.com/knr1997/assets-management-apiserver/internal/store"
)

type manufacturerKey string

const manufacturerCtx manufacturerKey = "manufacturer"

func getManufacturerFromCtx(r *http.Request) *store.Manufacturer {
	manufacturer, _ := r.Context().Value(manufacturerCtx).(*store.Manufacturer)
	return manufacturer
}

func (app *application) manufacturerContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "manufacturerID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		manufacturer, err := app.store.Manufacturer.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, manufacturerCtx, manufacturer)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// CreateManufacturer godoc
//
//	@Summary		Creates a manufacturer
//	@Description	Creates a manufacturer
//	@Tags			manufacturers
//	@Accept			json
//	@Produce		json
//	@Param			payload	body		CreateManufacturerPayload	true	"Manufacturer payload"
//	@Success		201		{object}	store.Manufacturer
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/manufacturers [manufacturer]
func (app *application) createManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	var payload requests.CreateManufacturerPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	manufacturer := &store.Manufacturer{
		Name:  payload.Name,
		Email: payload.Email,
	}

	ctx := r.Context()

	if err := app.store.Manufacturer.Create(ctx, manufacturer); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, manufacturer); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) updateManufacturer(ctx context.Context, manufacturer *store.Manufacturer) error {
	if err := app.store.Manufacturer.Update(ctx, manufacturer); err != nil {
		return err
	}

	return nil
}

// UpdateManufacturer godoc
//
//	@Summary		Updates a manufacturer
//	@Description	Updates a manufacturer by ID
//	@Tags			manufacturers
//	@Accept			json
//	@Produce		json
//	@Param			id		path		int					true	"Manufacturer ID"
//	@Param			payload	body		UpdateManufacturerPayload	true	"Manufacturer payload"
//	@Success		200		{object}	store.Manufacturer
//	@Failure		400		{object}	error
//	@Failure		401		{object}	error
//	@Failure		404		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/manufacturers/{id} [patch]
func (app *application) updateManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	manufacturer := getManufacturerFromCtx(r)

	var payload requests.UpdateManufacturerPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name != nil {
		manufacturer.Name = *payload.Name
	}
	if payload.Email != nil {
		manufacturer.Email = *payload.Email
	}

	ctx := r.Context()

	if err := app.updateManufacturer(ctx, manufacturer); err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, manufacturer); err != nil {
		app.internalServerError(w, r, err)
	}
}

// GetManufacturer godoc
//
//	@Summary		Fetches a manufacturer
//	@Description	Fetches a manufacturer by ID
//	@Tags			manufacturers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Manufacturer ID"
//	@Success		200	{object}	store.Manufacturer
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/manufacturers/{id} [get]
func (app *application) getManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	manufacturer := getManufacturerFromCtx(r)

	response := responses.NewManufacturerResponse(manufacturer)

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAllManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	Manufacturers, err := app.store.Manufacturer.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := responses.NewManufacturersResponse(Manufacturers)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DeleteManufacturer godoc
//
//	@Summary		Deletes a manufacturer
//	@Description	Delete a manufacturer by ID
//	@Tags			manufacturers
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Manufacturer ID"
//	@Success		204	{object} string
//	@Failure		404	{object}	error
//	@Failure		500	{object}	error
//	@Security		ApiKeyAuth
//	@Router			/manufacturers/{id} [delete]
func (app *application) deleteManufacturerHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "ManufacturerID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Manufacturer.Delete(ctx, id); err != nil {
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
