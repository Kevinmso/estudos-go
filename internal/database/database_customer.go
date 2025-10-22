package database

import (
	"context"

	"github.com/kevinmso/estudos-go/internal/models"
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
