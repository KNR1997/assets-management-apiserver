package store

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

var (
	ErrNotFound          = errors.New("resource not found")
	ErrConflict          = errors.New("resource already exists")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Users UsersStore
	Roles interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *gorm.DB) Storage {
	return Storage{
		Users: UsersStore{db},
		Roles: &RoleStore{db},
	}
}
