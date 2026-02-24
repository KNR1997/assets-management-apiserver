package store

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type AssetStatus string

const (
	AssetAvailable AssetStatus = "AVAILABLE"
	AssetAssigned  AssetStatus = "ASSIGNED"
	AssetRepair    AssetStatus = "REPAIR"
	AssetRetired   AssetStatus = "RETIRED"

	AssetPending       AssetStatus = "PENDING"
	AssetReadyToDeploy AssetStatus = "READY_TO_DEPLOY"
	AssetArchived      AssetStatus = "ARCHIVED"
	AssetBroken        AssetStatus = "BROKEN"
	AssetLostStolen    AssetStatus = "LOST_STOLEN"
)

type Asset struct {
	ID           int64  `gorm:"primaryKey"`
	Name         string `gorm:"size:150;not null"`
	SerialNumber string `gorm:"size:100;uniqueIndex;not null"`
	Tag          string `gorm:"size:100;uniqueIndex;not null"`
	Description  string `gorm:"size:255"`

	ModelID int64 `gorm:"not null;index"`
	Model   Model `gorm:"constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`

	Status AssetStatus `gorm:"type:varchar(20);not null;default:'AVAILABLE'"`

	PurchaseDate time.Time
	PurchaseCost float64

	UsefulLifeYears int     // for depreciation
	SalvageValue    float64 // optional

	Location string `gorm:"size:100"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type AssetStore struct {
	db *gorm.DB
}

func (s *AssetStore) GetAll(ctx context.Context) ([]Asset, error) {
	var assets []Asset

	err := s.db.WithContext(ctx).Preload("Model").Find(&assets).Error
	if err != nil {
		return nil, err
	}

	return assets, nil
}

func (s *AssetStore) GetByID(ctx context.Context, id int64) (*Asset, error) {
	var asset Asset

	err := s.db.WithContext(ctx).Preload("Model").First(&asset, id).Error
	if err != nil {
		return nil, err
	}

	return &asset, nil
}

func (s AssetStore) Create(ctx context.Context, asset *Asset) error {
	return s.db.WithContext(ctx).Create(asset).Error
}

func (s *AssetStore) Update(ctx context.Context, asset *Asset) error {
	result := s.db.WithContext(ctx).
		Model(&Asset{}).
		Where("id = ?", asset.ID).
		Updates(map[string]interface{}{
			"name":         asset.Name,
			"tag":          asset.Tag,
			"serialNumber": asset.SerialNumber,
			"description":  asset.Description,
		})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *AssetStore) Delete(ctx context.Context, id int64) error {
	result := s.db.WithContext(ctx).
		Delete(&Asset{}, id)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return ErrNotFound
	}

	return nil
}

func (s *AssetStore) UpdateStatus(ctx context.Context, assetID int64, status AssetStatus) error {
	result := s.db.WithContext(ctx).
		Model(&Asset{}).
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
