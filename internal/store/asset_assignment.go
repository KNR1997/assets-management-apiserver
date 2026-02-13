package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AssetAssignment struct {
	ID int64 `gorm:"primaryKey"`

	AssetID int64 `gorm:"not null;index"`
	Asset   Asset `gorm:"constraint:OnDelete:CASCADE;"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:RESTRICT;"`

	AssignedAt time.Time  `gorm:"not null"`
	ReturnedAt *time.Time // nullable

	Notes string `gorm:"size:255"`

	CreatedAt time.Time
}

type AssetAssignmentStore struct {
	db *gorm.DB
}

func (s AssetAssignmentStore) Create(ctx context.Context, asset *AssetAssignment) error {
	return s.db.WithContext(ctx).Create(asset).Error
}
