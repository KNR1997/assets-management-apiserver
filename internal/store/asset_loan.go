package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AssetLoan struct {
	ID        int64  `gorm:"primaryKey"`
	AssetName string `gorm:"size:100;uniqueIndex;not null"`

	AssetID int64 `gorm:"not null;index"`
	Asset   Asset `gorm:"constraint:OnDelete:CASCADE;"`

	UserID int64 `gorm:"not null;index"`
	User   User  `gorm:"constraint:OnDelete:RESTRICT;"`

	CheckoutDate        time.Time  `gorm:"not null"`
	ExpectedCheckinDate *time.Time // nullable
	ActualReturnDate    *time.Time

	Status AssetStatus `gorm:"type:varchar(20);not null;default:'AVAILABLE'"`

	Notes string `gorm:"size:255"`

	CreatedAt time.Time
}

type AssetLoanStore struct {
	db *gorm.DB
}

func (s AssetLoanStore) Create(ctx context.Context, assetLoan *AssetLoan) error {
	return s.db.WithContext(ctx).Create(assetLoan).Error
}

func (s *AssetLoanStore) UpdateStatus(ctx context.Context, assetID int64, status AssetStatus) error {
	result := s.db.WithContext(ctx).
		Model(&AssetLoan{}).
		Where("id = ?", assetID).
		Update("status", status)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}
