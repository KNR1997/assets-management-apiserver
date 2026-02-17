package requests

type CreateManufacturerPayload struct {
	Name  string `json:"name" validate:"required,max=100"`
	Email string `json:"email" validate:"required,max=100"`
}

type UpdateManufacturerPayload struct {
	Name  *string `json:"name" validate:"required,max=100"`
	Email *string `json:"email" validate:"required,max=100"`
}
