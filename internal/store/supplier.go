package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Supplier struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"size:100;uniqueIndex;not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type SupplierStore struct {
	db *gorm.DB
}

func (s *SupplierStore) GetAll(ctx context.Context) ([]Supplier, error) {
	var categories []Supplier

	err := s.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *SupplierStore) GetByID(ctx context.Context, id int64) (*Supplier, error) {
	var Supplier Supplier

	err := s.db.WithContext(ctx).
		First(&Supplier, id).
		Error
	if err != nil {
		return nil, err
	}

	return &Supplier, nil
}

func (s SupplierStore) Create(ctx context.Context, supplier *Supplier) error {
	return s.db.WithContext(ctx).Create(supplier).Error
}

func (s *SupplierStore) Update(ctx context.Context, supplier *Supplier) error {
	result := s.db.WithContext(ctx).
		Model(&Supplier{}).
		Where("id = ?", supplier.ID).
		Updates(map[string]interface{}{
			"name": supplier.Name,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *SupplierStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Supplier{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
