package services

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
)

type KonsumenService interface {
	Create(req dto.CreateKonsumenRequest, ctx *gin.Context) (*models.Konsumen, error)
	FindAll(ctx *gin.Context) ([]models.Konsumen, error)
	FindById(idStr string) (models.Konsumen, error)
	Update(idStr string, req dto.UpdateKonsumenRequest) (*models.Konsumen, error)
}