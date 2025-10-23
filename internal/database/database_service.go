package database

import (
	"context"

	"github.com/kevinmso/estudos-go/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Service, error) {
	var services []models.Service
	result := c.DB.WithContext(ctx).
		Find(&services)

	if result.Error != nil {
		return nil, result.Error
	}
	return services, nil
}
