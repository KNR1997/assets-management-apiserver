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

type departmentKey string

const departmentCtx departmentKey = "department"

func getdepartmentFromCtx(r *http.Request) *store.Department {
	department, _ := r.Context().Value(departmentCtx).(*store.Department)
	return department
}

func (app *application) departmentContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "departmentID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			app.internalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		department, err := app.store.Department.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				app.notFoundResponse(w, r, err)
			default:
				app.internalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, departmentCtx, department)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type CreatedepartmentPayload struct {
	Name  string `json:"name" validate:"required,max=100"`
	Notes string `json:"description"`
}

func (app *application) createdepartmentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreatedepartmentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	post := &store.Department{
		Name:  payload.Name,
		Notes: payload.Notes,
	}

	ctx := r.Context()

	if err := app.store.Department.Create(ctx, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	if err := app.jsonResponse(w, http.StatusCreated, post); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type UpdatedepartmentPayload struct {
	Name  *string `json:"name" validate:"omitempty,max=100"`
	Notes *string `json:"notes"`
}

func (app *application) updatedepartment(ctx context.Context, department *store.Department) error {
	if err := app.store.Department.Update(ctx, department); err != nil {
		return err
	}

	return nil
}

func (app *application) updatedepartmentHandler(w http.ResponseWriter, r *http.Request) {
	department := getdepartmentFromCtx(r)

	var payload UpdatedepartmentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Name != nil {
		department.Name = *payload.Name
	}
	if payload.Notes != nil {
		department.Notes = *payload.Notes
	}

	ctx := r.Context()

	if err := app.updatedepartment(ctx, department); err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, department); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) getdepartmentHandler(w http.ResponseWriter, r *http.Request) {
	department := getdepartmentFromCtx(r)

	response := TodepartmentResponse(department)

	if err := app.jsonResponse(w, http.StatusOK, response); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}

type departmentResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Notes string `json:"notes"`
}

func TodepartmentResponse(a *store.Department) departmentResponse {
	return departmentResponse{
		ID:    a.ID,
		Name:  a.Name,
		Notes: a.Notes,
	}
}

func TodepartmentResponseList(categories []store.Department) []departmentResponse {
	responses := make([]departmentResponse, 0, len(categories))

	for i := range categories {
		responses = append(responses, TodepartmentResponse(&categories[i]))
	}

	return responses
}

func (app *application) getAlldepartmentHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	categories, err := app.store.Department.GetAll(ctx)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	response := TodepartmentResponseList(categories)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (app *application) deletedepartmentHandler(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "departmentID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	ctx := r.Context()

	if err := app.store.Department.Delete(ctx, id); err != nil {
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
