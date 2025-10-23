package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kevinmso/estudos-go/internal/dberrors"
	"github.com/kevinmso/estudos-go/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetProductsByVendor(ctx context.Context, vendorId string) ([]models.Product, error) {
	var products []models.Product
	result := c.DB.WithContext(ctx).
		Where(&models.Product{VendorId: vendorId}).Find(&products)

	if result.Error != nil {
		return nil, result.Error
	}
	return products, nil
}

func (c Client) AddProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	product.ProductId = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&product)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}
