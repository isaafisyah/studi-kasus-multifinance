package repositories

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"gorm.io/gorm"
)

type KonsumenRepositoryImpl struct {
	db *gorm.DB
}

func NewKonsumenRepository(db *gorm.DB) KonsumenRepository {
	return &KonsumenRepositoryImpl{db: db}
}

func (k *KonsumenRepositoryImpl) FindById(id int) (models.Konsumen, error) {
	var konsumen models.Konsumen
	err := k.db.First(&konsumen, id).Error
	return konsumen, err
}

func (k *KonsumenRepositoryImpl) FindAll(ctx *gin.Context) ([]models.Konsumen, error) {
	var konsumens []models.Konsumen
	if err := k.db.
		Scopes(Paginate(ctx)).
		Find(&konsumens).Error; err != nil {
		return nil, err
	}
	return konsumens, nil
}

func (k *KonsumenRepositoryImpl) Save(konsumen models.Konsumen) (error) {
	err := k.db.Save(&konsumen).Error
	return err
}

func (k *KonsumenRepositoryImpl) Update(konsumen models.Konsumen) (models.Konsumen, error) {
	err := k.db.Save(&konsumen).Error
	return konsumen, err
}

func (k *KonsumenRepositoryImpl) Delete(konsumen models.Konsumen) (error) {
	err := k.db.Delete(&konsumen).Error
	return err
}

func Paginate(r *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.Request.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		sort := q.Get("sort")
		if sort == "" {
			sort = "id"
		}

		direction := q.Get("direction")
		if direction != "asc" && direction != "desc" {
			direction = "desc"
		}

		offset := (page - 1) * pageSize
		query := db.Offset(offset).Limit(pageSize).Order(sort + " " + direction)

		return query
	}
}

func (k *KonsumenRepositoryImpl) FindByNIK(nik string) (models.Konsumen, error) {
	var konsumen models.Konsumen
	err := k.db.Where("nik = ?", nik).First(&konsumen).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return models.Konsumen{}, nil
	}
	if err != nil {
		return models.Konsumen{}, err
	}
	return konsumen, err
	
}

