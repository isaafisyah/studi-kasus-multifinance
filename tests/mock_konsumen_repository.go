package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockKonsumenRepository struct {
	mock.Mock
}

func (m *MockKonsumenRepository) FindAll(ctx *gin.Context) ([]models.Konsumen, error) {
	args := m.Called()
	return args.Get(0).([]models.Konsumen), args.Error(1)
}

func (m *MockKonsumenRepository) FindById(id int) (models.Konsumen, error) {
	args := m.Called(id)
	return args.Get(0).(models.Konsumen), args.Error(1)
}

func (m *MockKonsumenRepository) Save(konsumen models.Konsumen) (error) {
	args := m.Called(konsumen)
	return args.Error(0)
}

func (m *MockKonsumenRepository) Update(konsumen models.Konsumen) (models.Konsumen, error) {
	args := m.Called(konsumen)
	return args.Get(0).(models.Konsumen), args.Error(1)
}

func (m *MockKonsumenRepository) Delete(konsumen models.Konsumen) (error) {
	args := m.Called(konsumen)
	return args.Error(0)
}

func (m *MockKonsumenRepository) FindByNIK(nik string) (models.Konsumen, error) {
	args := m.Called(nik)
	return args.Get(0).(models.Konsumen), args.Error(1)
}