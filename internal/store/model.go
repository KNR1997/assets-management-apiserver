package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID   int64  `gorm:"primaryKey"`
	Name string `gorm:"size:100;uniqueIndex;not null"`

	CategoryID int64    `gorm:"not null;index"`
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	ManufacturerID int64        `gorm:"not null;index"`
	Manufacturer   Manufacturer `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	ModelNumber string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type ModelStore struct {
	db *gorm.DB
}

func (s *ModelStore) GetAll(ctx context.Context) ([]Model, error) {
	var models []Model

	err := s.db.WithContext(ctx).Find(&models).Error
	if err != nil {
		return nil, err
	}

	return models, nil
}

func (s *ModelStore) GetByID(ctx context.Context, id int64) (*Model, error) {
	var model Model

	err := s.db.WithContext(ctx).
		First(&model, id).
		Error
	if err != nil {
		return nil, err
	}

	return &model, nil
}

func (s ModelStore) Create(ctx context.Context, model *Model) error {
	return s.db.WithContext(ctx).Create(model).Error
}

func (s *ModelStore) Update(ctx context.Context, model *Model) error {
	result := s.db.WithContext(ctx).
		Model(&Model{}).
		Where("id = ?", model.ID).
		Updates(map[string]interface{}{
			"name": model.Name,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *ModelStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Model{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
