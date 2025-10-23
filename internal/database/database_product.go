package database

import (
	"context"

	"github.com/kevinmso/estudos-go/internal/models"
)

func (c Client) GetProductsByVendor(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).Find(&products).
		Where(&models.Product{VendorId: vendorId})

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}
