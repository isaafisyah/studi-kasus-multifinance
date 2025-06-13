package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/log"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
)

type KonsumenControllerImpl struct {
	konsumenService services.KonsumenService
}

func NewKonsumenController(konsumenService services.KonsumenService) KonsumenController {
	return &KonsumenControllerImpl{konsumenService}
}

func (c *KonsumenControllerImpl) FindAll(ctx *gin.Context)  {
	log.GetLogger("KonsumenController").Info("All Konsumen Start")
	konsumens, err := c.konsumenService.FindAll(ctx)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"data": konsumens})
}

func (c *KonsumenControllerImpl) FindById(ctx *gin.Context)  {
	log.GetLogger("KonsumenController").Info("Detail Konsumen Start")
	idStr := ctx.Param("id")
	konsumen, err := c.konsumenService.FindById(idStr)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.GetLogger("KonsumenController").Info("Detail Konsumen End")
	ctx.JSON(200, gin.H{"data": konsumen})
}

func (c *KonsumenControllerImpl) Create(ctx *gin.Context)  {
	log.GetLogger("KonsumenController").Info("Create Konsumen Start")
	var req dto.CreateKonsumenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := c.konsumenService.Create(req, ctx)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.GetLogger("KonsumenController").Info("Create Konsumen End")
	ctx.JSON(http.StatusCreated, gin.H{"message" : "data berhasil disimpan"})
}

func (c *KonsumenControllerImpl) Update(ctx *gin.Context)  {
	log.GetLogger("KonsumenController").Info("Update Konsumen Start")
	var req dto.UpdateKonsumenRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("id")
	data, err := c.konsumenService.Update(id , req)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	log.GetLogger("KonsumenController").Info("Update Konsumen End")
	ctx.JSON(200, gin.H{"data": data})
}

func (c *KonsumenControllerImpl) Delete(ctx *gin.Context)  {
	log.GetLogger("KonsumenController").Info("Delete Konsumen Start")
	idParam := ctx.Param("id")

	if err := c.konsumenService.Delete(idParam); err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}
	log.GetLogger("KonsumenController").Info("Delete Konsumen End")
	ctx.JSON(200, gin.H{"message": "data berhasil dihapus"})
}

