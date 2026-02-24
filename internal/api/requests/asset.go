package requests

import "time"

type CreateAssetPayload struct {
	Name         string `json:"name" validate:"required,max=100"`
	Tag          string `json:"tag" validate:"required,max=100"`
	SerialNumber string `json:"serialNumber" validate:"required,max=100"`
	Description  string `json:"description"`
	ModelID      int64  `json:"modelId" validate:"required"`
	Status       string `json:"status" validate:"required"`
	Notes        string `json:"notes"`
}

type UpdateAssetPayload struct {
	Name         *string `json:"name" validate:"omitempty,max=100"`
	Tag          *string `json:"tag" validate:"required,max=100"`
	SerialNumber *string `json:"serialNumber" validate:"required,max=100"`
	Description  *string `json:"description"`
	ModelID      *int64  `json:"modelId" validate:"required"`
	Status       *string `json:"status" validate:"required"`
	Notes        *string `json:"notes"`
}

type CheckoutAssetPayload struct {
	AssetName           string     `json:"assetName" validate:"required,max=100"`
	AssetID             int64      `json:"assetId" validate:"required"`
	UserID              int64      `json:"userId" validate:"required"`
	CheckoutDate        time.Time  `json:"checkoutDate" validate:"required"`
	ExpectedCheckinDate *time.Time `json:"expectedCheckinDate"`
	Notes               string     `json:"notes"`
}

type CheckinAssetPayload struct {
	AssetName   string    `json:"assetName" validate:"required,max=100"`
	AssetID     int64     `json:"assetId" validate:"required"`
	CheckinDate time.Time `json:"checkinDate" validate:"required"`
	Status      string    `json:"status"`
	Notes       string    `json:"notes"`
}
