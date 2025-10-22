package database

import (
	"context"

	"github.com/kevinmso/estudos-go/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendor, error) {
	var vendors []models.Vendor
	result := c.DB.WithContext(ctx).Find(&vendors)

	if result.Error != nil {
		return nil, result.Error
	}
	return vendors, nil
}
