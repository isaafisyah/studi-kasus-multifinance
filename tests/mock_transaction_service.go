package tests

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockRecordTransactionService struct {
	mock.Mock
}

func (m *MockRecordTransactionService) Create(req dto.CreateRecordTransactionRequest) (*models.RecordTransaction, error) {
	args := m.Called(req)
	return args.Get(0).(*models.RecordTransaction), args.Error(1)
}

func (m *MockRecordTransactionService) FindAll() ([]models.RecordTransaction, error) {
	args := m.Called()
	return args.Get(0).([]models.RecordTransaction), args.Error(1)
	
}

func (m *MockRecordTransactionService) FindById(id string) (models.RecordTransaction, error) {
	args := m.Called(id)
	return args.Get(0).(models.RecordTransaction), args.Error(1)
}