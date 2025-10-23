package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kevinmso/estudos-go/internal/dberrors"
	"github.com/kevinmso/estudos-go/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).Find(&vendors)

	if result.Error != nil {
		return nil, result.Error
	}
	return vendors, nil
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendor) (*models.Vendor, error) {
	vendor.VendorId = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&vendor)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}
