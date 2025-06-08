package services

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
)

type LimitService interface {
	FindAll() ([]models.Limit, error)
	FindById(id string) (models.Limit, error)
	Create(req dto.CreateLimitRequest) (*models.Limit, error)
	Update(id string, req dto.UpdateLimitRequest) (*models.Limit, error)
}