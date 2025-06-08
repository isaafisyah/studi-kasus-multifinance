package services

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
)

type RecordTransactionService interface {
	FindAll() ([]models.RecordTransaction, error)
	FindById(id string) (models.RecordTransaction, error)
	Create(req dto.CreateRecordTransactionRequest) (*models.RecordTransaction, error)
}