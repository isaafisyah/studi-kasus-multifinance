package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
	"gorm.io/gorm"
)

func InitializeRoutesLimit(db *gorm.DB, r *gin.Engine)  {
	repo := repositories.NewLimitRepository(db)
	service := services.NewLimitService(repo)

	controller := controllers.NewLimitController(service, repo)

	r.GET("/limit", controller.FindAll)
	r.GET("/limit/:id", controller.FindById)
	r.POST("/limit", controller.Create)
	r.PUT("/limit/:id", controller.Update)
	r.DELETE("/limit/:id", controller.Delete)	
}