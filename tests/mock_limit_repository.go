package tests

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockLimitRepository struct {
	mock.Mock
}

func (m *MockLimitRepository) FindAll() ([]models.Limit, error) {
	args := m.Called()
	return args.Get(0).([]models.Limit), args.Error(1)
}

func (m *MockLimitRepository) FindById(id int) (models.Limit, error) {
	args := m.Called(id)
	return args.Get(0).(models.Limit), args.Error(1)
}

func (m *MockLimitRepository) Save(limit models.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func (m *MockLimitRepository) Update(limit models.Limit) (models.Limit, error) {
	args := m.Called(limit)
	return args.Get(0).(models.Limit), args.Error(1)
}

func (m *MockLimitRepository) Delete(limit models.Limit) error {
	args := m.Called(limit)
	return args.Error(0)
}

func (m *MockLimitRepository) FindByKonsumenTenor(konsumenID int, tenor uint8) (models.Limit, error) {
	args := m.Called(konsumenID, tenor)
	return args.Get(0).(models.Limit), args.Error(1)
}
