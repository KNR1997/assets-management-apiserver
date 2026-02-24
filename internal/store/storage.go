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
	Users           UsersStore
	Manufacturer    ManufacturerStore
	Category        CategoryStore
	Asset           AssetStore
	AssetAssignment AssetAssignmentStore
	AssetLoan       AssetLoanStore
	AssetLog        AssetLogStore
	Model           ModelStore
	Department      DepartmentStore
	Supplier        SupplierStore
	Roles           interface {
		GetByName(context.Context, string) (*Role, error)
	}
}

func NewStorage(db *gorm.DB, auditService AuditService) Storage {
	return Storage{
		Users:           UsersStore{db},
		Manufacturer:    ManufacturerStore{db},
		Category:        CategoryStore{db, auditService},
		Asset:           AssetStore{db},
		AssetAssignment: AssetAssignmentStore{db},
		AssetLoan:       AssetLoanStore{db},
		AssetLog:        AssetLogStore{db},
		Model:           ModelStore{db},
		Department:      DepartmentStore{db},
		Supplier:        SupplierStore{db},
		Roles:           &RoleStore{db},
	}
}
