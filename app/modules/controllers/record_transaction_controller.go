package controllers

import "github.com/gin-gonic/gin"

type RecordTransactionController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
}