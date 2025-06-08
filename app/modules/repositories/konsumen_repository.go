package repositories

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
)

type KonsumenRepository interface{
	FindAll(ctx *gin.Context) ([]models.Konsumen, error)
	FindById(id int) (models.Konsumen, error)
	Save(konsumen models.Konsumen) (error)
	Update(konsumen models.Konsumen) (models.Konsumen, error)
	Delete(konsumen models.Konsumen) (error)
	FindByNIK(nik string) (models.Konsumen, error)
}