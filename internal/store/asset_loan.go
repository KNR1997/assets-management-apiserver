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

	Notes string `gorm:"size:255"`

	CreatedAt time.Time
}

type AssetLoanStore struct {
	db *gorm.DB
}

func (s AssetLoanStore) Create(ctx context.Context, assetLoan *AssetLoan) error {
	return s.db.WithContext(ctx).Create(assetLoan).Error
}
