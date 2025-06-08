package tests

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockLimitService struct {
	mock.Mock
}

func (m *MockLimitService) FindAll() ([]models.Limit, error) {
	args := m.Called()
	return args.Get(0).([]models.Limit), args.Error(1)
}

func (m *MockLimitService) FindById(id string) (models.Limit, error) {
	args := m.Called(id)
	return args.Get(0).(models.Limit), args.Error(1)
}

func (m *MockLimitService) Create(req dto.CreateLimitRequest) (*models.Limit, error) {
	args := m.Called(req)
	return args.Get(0).(*models.Limit), args.Error(1)
}

func (m *MockLimitService) Update(id string, req dto.UpdateLimitRequest) (*models.Limit, error) {
	args := m.Called(id, req)
	return args.Get(0).(*models.Limit), args.Error(1)
}