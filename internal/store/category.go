package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID          int64  `gorm:"primaryKey"`
	Name        string `gorm:"size:100;uniqueIndex;not null"`
	Description string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type CategoryStore struct {
	db           *gorm.DB
	auditService AuditService
}

func (s *CategoryStore) GetAll(ctx context.Context) ([]Category, error) {
	var categories []Category

	err := s.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryStore) GetByID(ctx context.Context, id int64) (*Category, error) {
	var category Category

	err := s.db.WithContext(ctx).
		First(&category, id).
		Error
	if err != nil {
		return nil, err
	}

	return &category, nil
}

func (s CategoryStore) Create(ctx context.Context, category *Category) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(category).Error; err != nil {
			return err
		}

		return s.auditService.LogCreate(
			ctx,
			tx,
			"categories",
			category.ID,
			category,
		)
	})
}

func (s *CategoryStore) Update(ctx context.Context, category *Category) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var oldCategory Category
		if err := tx.First(&oldCategory, category.ID).Error; err != nil {
			return err
		}

		if err := tx.Save(category).Error; err != nil {
			return err
		}

		return s.auditService.LogUpdate(
			ctx,
			tx,
			"categories",
			category.ID,
			oldCategory,
			category,
		)
	})
}

func (s *CategoryStore) Delete(ctx context.Context, id int64) error {
	return s.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		var category Category
		if err := tx.First(&category, id).Error; err != nil {
			return err
		}

		if err := tx.Delete(&category, id).Error; err != nil {
			return err
		}

		return s.auditService.LogDelete(
			ctx,
			tx,
			"categories",
			id,
			category,
		)
	})
}
