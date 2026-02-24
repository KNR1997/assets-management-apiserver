package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Department struct {
	ID    int64  `gorm:"primaryKey"`
	Name  string `gorm:"size:100;uniqueIndex;not null"`
	Notes string `gorm:"size:255"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

type DepartmentStore struct {
	db *gorm.DB
}

func (s *DepartmentStore) GetAll(ctx context.Context) ([]Department, error) {
	var categories []Department

	err := s.db.WithContext(ctx).Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *DepartmentStore) GetByID(ctx context.Context, id int64) (*Department, error) {
	var Department Department

	err := s.db.WithContext(ctx).
		First(&Department, id).
		Error
	if err != nil {
		return nil, err
	}

	return &Department, nil
}

func (s DepartmentStore) Create(ctx context.Context, department *Department) error {
	return s.db.WithContext(ctx).Create(department).Error
}

func (s *DepartmentStore) Update(ctx context.Context, department *Department) error {
	result := s.db.WithContext(ctx).
		Model(&Department{}).
		Where("id = ?", department.ID).
		Updates(map[string]interface{}{
			"name":  department.Name,
			"notes": department.Notes,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *DepartmentStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Department{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
