package store

import (
	"context"

	"gorm.io/gorm"
)

type Manufacturer struct {
	ID    int64  `json:"id"`
	Name  string `gorm:"uniqueIndex;not null" json:"name"`
	Email string `gorm:"uniqueIndex;not null" json:"email"`
}

type ManufacturerStore struct {
	db *gorm.DB
}

func (s *ManufacturerStore) GetAll(ctx context.Context) ([]Manufacturer, error) {
	var manufacturers []Manufacturer

	err := s.db.WithContext(ctx).Find(&manufacturers).Error
	if err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func (s *ManufacturerStore) GetByID(ctx context.Context, id int64) (*Manufacturer, error) {
	var manufacturer Manufacturer

	err := s.db.WithContext(ctx).
		First(&manufacturer, id).
		Error
	if err != nil {
		return nil, err
	}

	return &manufacturer, nil
}

func (s ManufacturerStore) Create(ctx context.Context, asset *Manufacturer) error {
	return s.db.WithContext(ctx).Create(asset).Error
}

func (s *ManufacturerStore) Update(ctx context.Context, manufacturer *Manufacturer) error {
	result := s.db.WithContext(ctx).
		Model(&Manufacturer{}).
		Where("id = ?", manufacturer.ID).
		Updates(map[string]interface{}{
			"name":  manufacturer.Name,
			"email": manufacturer.Email,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *ManufacturerStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Manufacturer{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
