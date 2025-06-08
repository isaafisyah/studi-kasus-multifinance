package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
)

type KonsumenControllerImpl struct {
	konsumenService services.KonsumenService
	repo repositories.KonsumenRepository
}

func NewKonsumenController(konsumenService services.KonsumenService, repo repositories.KonsumenRepository) KonsumenController {
	return &KonsumenControllerImpl{konsumenService, repo}
}

func (c *KonsumenControllerImpl) FindAll(ctx *gin.Context)  {
	konsumens, err := c.konsumenService.FindAll(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": konsumens})
}

func (c *KonsumenControllerImpl) FindById(ctx *gin.Context)  {
	idStr := ctx.Param("id")
	konsumen, err := c.konsumenService.FindById(idStr)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": konsumen})
}

func (c *KonsumenControllerImpl) Create(ctx *gin.Context)  {
	var req dto.CreateKonsumenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	konsumen, err := c.konsumenService.Create(req, ctx)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.Save(*konsumen); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message" : "data berhasil disimpan"})
}

func (c *KonsumenControllerImpl) Update(ctx *gin.Context)  {
	var req dto.UpdateKonsumenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	konsumen, err := c.konsumenService.Update(id , req)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	data, err := c.repo.Update(*konsumen)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": data})
}

func (c *KonsumenControllerImpl) Delete(ctx *gin.Context)  {
	idParam := ctx.Param("id")
	konsumen, err := c.konsumenService.FindById(idParam)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := c.repo.Delete(konsumen); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "data berhasil dihapus"})
}

