package main

import (
	"net/http"

	"github.com/knr1997/assets-management-apiserver/internal/store"
)

type CreateAssetAssignmentPayload struct {
	AssetID int64 `json:"assetID" validate:"required"`
	UserID  int64 `json:"userID" validate:"required"`
}

func (app *application) CreateAssetAssignmentHandler(w http.ResponseWriter, r *http.Request) {
	var payload CreateAssetAssignmentPayload
	if err := readJSON(w, r, &payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if err := Validate.Struct(payload); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	asset_assignment := &store.AssetAssignment{
		AssetID: payload.AssetID,
		UserID:  payload.UserID,
	}

	ctx := r.Context()

	asset, err := app.store.Asset.GetByID(ctx, payload.AssetID)
	if err != nil {
		app.internalServerError(w, r, err)
		return
	}

	asset.Status = "ASSIGNED"

	if err := app.store.AssetAssignment.Create(ctx, asset_assignment); err != nil {
		app.internalServerError(w, r, err)
		return
	}

	app.updateAsset(ctx, asset)

	if err := app.jsonResponse(w, http.StatusCreated, asset_assignment); err != nil {
		app.internalServerError(w, r, err)
		return
	}
}
