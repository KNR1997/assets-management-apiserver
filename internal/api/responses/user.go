package responses

import "github.com/knr1997/assets-management-apiserver/internal/store"

type UserResponse struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

func NewUserResponse(u *store.User) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}
}

func NewUsersResponse(users []store.User) []UserResponse {
	responses := make([]UserResponse, len(users))

	for i := range users {
		responses[i] = NewUserResponse(&users[i])
	}

	return responses
}
