package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
	"gorm.io/gorm"
)

func InitializeRoutesKonsumen(db *gorm.DB, r *gin.Engine)  {
	konsumenRepository := repositories.NewKonsumenRepository(db)
	konsumenService := services.NewKonsumenService(konsumenRepository)
	konsumenController := controllers.NewKonsumenController(konsumenService)

	r.GET("/konsumen", konsumenController.FindAll)
	r.GET("/konsumen/:id", konsumenController.FindById)
	r.POST("/konsumen", konsumenController.Create)
	r.PUT("/konsumen/:id", konsumenController.Update)
	r.DELETE("/konsumen/:id", konsumenController.Delete)
}