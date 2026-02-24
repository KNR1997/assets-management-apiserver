package responses

import "github.com/knr1997/assets-management-apiserver/internal/store"

type ModelResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func NewModelResponse(u *store.Model) ModelResponse {
	return ModelResponse{
		ID:   u.ID,
		Name: u.Name,
	}
}

func NewModelsResponse(Models []store.Model) []ModelResponse {
	responses := make([]ModelResponse, len(Models))

	for i := range Models {
		responses[i] = NewModelResponse(&Models[i])
	}

	return responses
}
