package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kevinmso/estudos-go/internal/dberrors"
	"github.com/kevinmso/estudos-go/internal/models"
	"gorm.io/gorm"
)

func (c Client) GetCustomersByEmail(ctx context.Context, email string) ([]models.Customer, error) {
	var customers []models.Customer
	result := c.DB.WithContext(ctx).
		Where(models.Customer{Email: email}).Find(&customers)

	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	customer.CustomerId = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&customer)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}
