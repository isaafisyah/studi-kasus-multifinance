package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
)

type LimitControllerImpl struct {
	limitService services.LimitService
	repo repositories.LimitRepository
}

func NewLimitController(limitService services.LimitService, repo repositories.LimitRepository) LimitController {
	return &LimitControllerImpl{limitService, repo}
}

func (c *LimitControllerImpl) FindAll(ctx *gin.Context) {
	limit, err := c.limitService.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": limit})
}

func (c *LimitControllerImpl) FindById(ctx *gin.Context) {
	idStr := ctx.Param("id")
	limit, err := c.limitService.FindById(idStr)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": limit})
}

func (c *LimitControllerImpl) Create(ctx *gin.Context) {
	var req dto.CreateLimitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	limit, err := c.limitService.Create(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Save(*limit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "limit created successfully"})
}

func (c *LimitControllerImpl) Update(ctx *gin.Context) {
	var req dto.UpdateLimitRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("id")
	limit, err := c.limitService.Update(id, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _,err := c.repo.Update(*limit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "limit updated successfully"})
}

func (c *LimitControllerImpl) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	limit, err := c.limitService.FindById(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.repo.Delete(limit); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "limit deleted successfully"})
	
}



