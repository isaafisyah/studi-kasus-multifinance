package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/mock"
)

type MockKonsumenService struct {
	mock.Mock
}

func (m *MockKonsumenService) FindAll(ctx *gin.Context) ([]models.Konsumen, error) {
	args := m.Called()
	return args.Get(0).([]models.Konsumen), args.Error(1)
}

func (m *MockKonsumenService) FindById(id string) (models.Konsumen, error) {
	args := m.Called(id)
	return args.Get(0).(models.Konsumen), args.Error(1)
}

func (m *MockKonsumenService) Create(req dto.CreateKonsumenRequest, ctx *gin.Context) (*models.Konsumen, error) {
	args := m.Called(req, ctx)
	konsumen, ok := args.Get(0).(*models.Konsumen)
	if !ok {
		return nil, args.Error(1)
	}
	return konsumen, args.Error(1)
}

func (m *MockKonsumenService) Update(id string, req dto.UpdateKonsumenRequest) (*models.Konsumen, error) {
	args := m.Called(id, req)
	return args.Get(0).(*models.Konsumen), args.Error(1)
}
