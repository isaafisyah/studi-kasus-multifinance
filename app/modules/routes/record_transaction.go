package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
	"gorm.io/gorm"
)

func InitializeRoutesRecordTransaction(db *gorm.DB, r *gin.Engine)  {
	transactionRepository := repositories.NewRecordTransactionRepository(db)
	limitRepository := repositories.NewLimitRepository(db)
	transactionService := services.NewRecordTransactionService(transactionRepository, limitRepository)
	transactionController := controllers.NewRecordTransactionController(transactionService, transactionRepository)
	
	r.GET("/transaction", transactionController.FindAll)
	r.GET("/transaction/:id", transactionController.FindById)
	r.POST("/transaction", transactionController.Create)
}