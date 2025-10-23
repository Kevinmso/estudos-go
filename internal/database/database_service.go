package database

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/kevinmso/estudos-go/internal/dberrors"
	"github.com/kevinmso/estudos-go/internal/models"
	"gorm.io/gorm"
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

func (c Client) AddService(ctx context.Context, service *models.Service) (*models.Service, error) {
	service.ServiceId = uuid.NewString()
	result := c.DB.WithContext(ctx).
		Create(&service)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return service, nil
}
