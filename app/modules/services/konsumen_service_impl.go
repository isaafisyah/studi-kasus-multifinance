package services

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/log"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
)

type KonsumenServiceImpl struct {
	konsumenRepository repositories.KonsumenRepository
}

func NewKonsumenService(konsumenRepository repositories.KonsumenRepository) KonsumenService {
	return &KonsumenServiceImpl{konsumenRepository}
}

func (k *KonsumenServiceImpl) FindAll(ctx *gin.Context) ([]models.Konsumen, error) {
	return k.konsumenRepository.FindAll(ctx)
}

func (k *KonsumenServiceImpl) FindById(idStr string) (models.Konsumen, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.Konsumen{}, fmt.Errorf("invalid id")
	}
	return k.konsumenRepository.FindById(id)
}

func (k *KonsumenServiceImpl) Create(req dto.CreateKonsumenRequest, ctx *gin.Context) (error) {
	fileKtp, err := ctx.FormFile("foto_ktp")
	if err != nil {
		log.GetLogger("KonsumenService").Error("validation error: foto_ktp wajib diunggah")
		return fmt.Errorf("validation error: foto_ktp wajib diunggah")
	}
	// Simpan file ke disk
	filePathKtp := fmt.Sprintf("storage/uploads/ktp_%s.jpg", req.NIK)
	if err := ctx.SaveUploadedFile(fileKtp, filePathKtp); err != nil {
		log.GetLogger("KonsumenService").Error("validation error: gagal menyimpan foto KTP")
		return fmt.Errorf("validation error: gagal menyimpan foto KTP")
	}

	fileSelfie, err := ctx.FormFile("foto_selfie")
	if err != nil {
		log.GetLogger("KonsumenService").Error("validation error: foto_ktp wajib diunggah")
		return fmt.Errorf("validation error: foto_ktp wajib diunggah")
	}
	// Simpan file ke disk
	filePathSelfie := fmt.Sprintf("storage/uploads/selfie_%s.jpg", req.NIK)
	if err := ctx.SaveUploadedFile(fileSelfie, filePathSelfie); err != nil {
		log.GetLogger("KonsumenService").Error("validation error: gagal menyimpan foto KTP")
		return fmt.Errorf("validation error: gagal menyimpan foto KTP")
	}

	dir, _ := os.Getwd()
	fmt.Println("Current working dir:", dir)

	tanggalLahir, err := time.Parse("2006-01-02", req.TanggalLahir)
	if err != nil {
		log.GetLogger("KonsumenService").Error("validation error: format tanggal lahir salah")
		return fmt.Errorf("validation error: format tanggal lahir salah")
	}

	gaji, err := strconv.ParseInt(req.Gaji, 10, 64)
	if err != nil {
		log.GetLogger("KonsumenService").Error("validation error: gaji tidak valid")
		return fmt.Errorf("validation error: gaji tidak valid")
	}

	exist, _ := k.konsumenRepository.FindByNIK(req.NIK)
	if exist.NIK != "" {
		log.GetLogger("KonsumenService").Error("validation error: NIK sudah terdaftar")
		return fmt.Errorf("validation error: NIK sudah terdaftar")
	}

	konsumen := &models.Konsumen{
		NIK:          req.NIK,
		Fullname:     req.Fullname,
		LegalName:    req.LegalName,
		TempatLahir:  req.TempatLahir,
		TanggalLahir: tanggalLahir,
		Gaji:         gaji,
		FotoKtp:      filePathKtp,
		FotoSelfie:   filePathSelfie,
	}

	if err := k.konsumenRepository.Save(*konsumen); err != nil {
		log.GetLogger("KonsumenService").Error(err.Error())
		return fmt.Errorf("validation error: format tanggal lahir salah")
	}

	return nil
}

func (k *KonsumenServiceImpl) Update(idStr string, req dto.UpdateKonsumenRequest) (*models.Konsumen, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return nil, fmt.Errorf("invalid product ID")
	}

	konsumen, err := k.konsumenRepository.FindById(id)
	if err != nil {
		return nil, err
	}

	if req.Fullname != nil {
		konsumen.Fullname = *req.Fullname
	}
	if req.LegalName != nil {
		konsumen.LegalName = *req.LegalName
	}
	if req.TempatLahir != nil {
		konsumen.TempatLahir = *req.TempatLahir
	}
	if req.TanggalLahir != nil {
		t, err := time.Parse("2006-01-02", *req.TanggalLahir)
		if err != nil {
			return nil, fmt.Errorf("format tanggal lahir salah")
		}
		konsumen.TanggalLahir = t
	}
	if req.Gaji != nil {
		gaji, err := strconv.ParseInt(*req.Gaji, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("gaji tidak valid")
		}
		konsumen.Gaji = gaji
	}

	data, err := k.konsumenRepository.Update(konsumen)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		return nil, err
	}

	return &data, nil
}

func (k *KonsumenServiceImpl) Delete(idStr string) (error) {
	konsumen, err := k.FindById(idStr)
	if err != nil {
		log.GetLogger("KonsumenController").Error(err.Error())
		return err
	}
	return k.konsumenRepository.Delete(konsumen)
}
