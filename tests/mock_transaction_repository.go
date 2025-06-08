package tests

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockRecordTransactionRepository struct {
	mock.Mock
}

func (m *MockRecordTransactionRepository) FindAll() ([]models.RecordTransaction, error) {
	args := m.Called()
	return args.Get(0).([]models.RecordTransaction), args.Error(1)
}

func (m *MockRecordTransactionRepository) FindById(id int) (models.RecordTransaction, error) {
	args := m.Called(id)
	return args.Get(0).(models.RecordTransaction), args.Error(1)
}

func (m *MockRecordTransactionRepository) Save(record models.RecordTransaction) error {
	args := m.Called(record)
	return args.Error(0)
}

func (m *MockRecordTransactionRepository) FindByKonsumenID(id int) ([]models.RecordTransaction, error) {
	args := m.Called(id)
	return args.Get(0).([]models.RecordTransaction), args.Error(1)
}