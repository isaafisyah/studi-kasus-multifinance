package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/log"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
)

type RecordTransactionControllerImpl struct {
	recordTransactionService services.RecordTransactionService
	repo repositories.RecordTransactionRepository
}

func NewRecordTransactionController(recordTransactionService services.RecordTransactionService, repo repositories.RecordTransactionRepository) RecordTransactionController {
	return &RecordTransactionControllerImpl{recordTransactionService: recordTransactionService, repo: repo}
	
}

func (c *RecordTransactionControllerImpl) FindAll(ctx *gin.Context) {
	log.GetLogger("TransactionController").Info("All Transaction Start")
	recordTransactions, err := c.recordTransactionService.FindAll()
	if err != nil {
		log.GetLogger("TransactionController").Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.GetLogger("TransactionController").Info("All Transaction End")
	ctx.JSON(200, gin.H{"data": recordTransactions})
}

func (c *RecordTransactionControllerImpl) FindById(ctx *gin.Context) {
	log.GetLogger("TransactionController").Info("Detail Transaction Start")
	idStr := ctx.Param("id")
	recordTransaction, err := c.recordTransactionService.FindById(idStr)
	if err != nil {
		log.GetLogger("TransactionController").Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.GetLogger("TransactionController").Info("Detail Transaction End")
	ctx.JSON(200, gin.H{"data": recordTransaction})
}

func (c *RecordTransactionControllerImpl) Create(ctx *gin.Context) {
	log.GetLogger("TransactionController").Info("Create Transaction Start")

	var req dto.CreateRecordTransactionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.GetLogger("TransactionController").Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	recordTransaction, err := c.recordTransactionService.Create(req)
	if err != nil {
		log.GetLogger("TransactionController").Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Save(*recordTransaction); err != nil {
		log.GetLogger("TransactionController").Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	log.GetLogger("TransactionController").Info("Create Transaction End")
	ctx.JSON(http.StatusCreated, gin.H{"message": "record transaction created successfully"})
}