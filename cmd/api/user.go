package main

import (
	"context"
	"net/http"

	"github.com/knr1997/assets-management-apiserver/internal/api/responses"
	"github.com/knr1997/assets-management-apiserver/internal/store"
)

type userKey string

const userCtx userKey = "user"

func getUserFromContext(r *http.Request) *store.User {
	user, _ := r.Context().Value(userCtx).(*store.User)
	return user
}

type UpdateUserPayload struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
}

func (app *application) updateUserHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	var payload UpdateUserPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if payload.Username != nil {
		user.Username = *payload.Username
	}
	if payload.Email != nil {
		user.Email = *payload.Email
	}

	ctx := r.Context()

	if err := app.updateUser(ctx, user); err != nil {
		app.internalServerError(w, r, err)
	}

	if err := app.jsonResponse(w, http.StatusOK, user); err != nil {
		app.internalServerError(w, r, err)
	}
}

func (app *application) updateUser(ctx context.Context, user *store.User) error {
	if err := app.store.Users.Update(ctx, user); err != nil {
		return err
	}

	return nil
}

func (app *application) meDetailsHandler(w http.ResponseWriter, r *http.Request) {
	user := getUserFromContext(r)

	resp := responses.NewUserResponse(user)

	app.jsonResponse(w, http.StatusOK, resp)
}
