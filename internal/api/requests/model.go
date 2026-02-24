package requests

type CreateModelPayload struct {
	Name           string `json:"name" validate:"required,max=100"`
	CategoryID     int64  `json:"categoryID" validate:"required"`
	ManufacturerID int64  `json:"manufacturerID" validate:"required"`
	ModelNumber    string `json:"modelNumber" validate:"required,max=100"`
}

type UpdateModelPayload struct {
	Name           *string `json:"name" validate:"required,max=100"`
	CategoryID     int64   `json:"categoryID" validate:"required"`
	ManufacturerID int64   `json:"manufacturerID" validate:"required"`
	ModelNumber    string  `json:"modelNumber" validate:"required,max=100"`
}
