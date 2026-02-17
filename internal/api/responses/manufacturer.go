package responses

import "github.com/knr1997/assets-management-apiserver/internal/store"

type ManufacturerResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewManufacturerResponse(u *store.Manufacturer) ManufacturerResponse {
	return ManufacturerResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func NewManufacturersResponse(Manufacturers []store.Manufacturer) []ManufacturerResponse {
	responses := make([]ManufacturerResponse, len(Manufacturers))

	for i := range Manufacturers {
		responses[i] = NewManufacturerResponse(&Manufacturers[i])
	}

	return responses
}
