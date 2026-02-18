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

type modelKey string

const modelCtx categoryKey = "model"

func getModelFromCtx(r *http.Request) *store.Model {
	model, _ := r.Context().Value(modelCtx).(*store.Model)
	return model
}

func (app *application) modelContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "modelID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		model, err := app.store.Model.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, modelCtx, model)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) createModelHandler(w http.ResponseWriter, r *http.Request) {
	var payload requests.CreateModelPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	model := &store.Model{
		Name:           payload.Name,
		CategoryID:     payload.CategoryID,
		ManufacturerID: payload.ManufacturerID,
		ModelNumber:    payload.ModelNumber,
	}

	ctx := r.Context()

	if err := app.store.Model.Create(ctx, model); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, model); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) updateModel(ctx context.Context, Model *store.Model) error {
	if err := app.store.Model.Update(ctx, Model); err != nil {
		return err
	}

	return nil
}

func (app *application) updateModelHandler(w http.ResponseWriter, r *http.Request) {
	Model := getModelFromCtx(r)

	var payload requests.UpdateModelPayload

	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name != nil {
		Model.Name = *payload.Name
	}

	ctx := r.Context()

	if err := app.updateModel(ctx, Model); err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, Model); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getModelHandler(w http.ResponseWriter, r *http.Request) {
	Model := getModelFromCtx(r)

	response := responses.NewModelResponse(Model)

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

func (app *application) getAllModelHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	Models, err := app.store.Model.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := responses.NewModelsResponse(Models)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *application) deleteModelHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "ModelID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Model.Delete(ctx, id); err != nil {
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
