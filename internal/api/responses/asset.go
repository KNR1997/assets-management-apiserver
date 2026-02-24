package responses

import (
	"github.com/knr1997/assets-management-apiserver/internal/store"
)

type AssetResponse struct {
	ID           int64         `json:"id"`
	Name         string        `json:"name"`
	SerialNumber string        `json:"serialNumber"`
	Tag          string        `json:"tag"`
	Status       string        `json:"status"`
	Model        ModelResponse `json:"model"`
	Description  string        `json:"description"`
}

func NewAssetResponse(u *store.Asset) AssetResponse {
	return AssetResponse{
		ID:           u.ID,
		Name:         u.Name,
		SerialNumber: u.SerialNumber,
		Tag:          u.Tag,
		Status:       string(u.Status),
		Model:        NewModelResponse(&u.Model),
		Description:  u.Description,
	}
}

func NewAssetsResponse(assets []store.Asset) []AssetResponse {
	responses := make([]AssetResponse, len(assets))

	for i := range assets {
		responses[i] = NewAssetResponse(&assets[i])
	}

	return responses
}
