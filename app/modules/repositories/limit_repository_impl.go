package repositories

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"gorm.io/gorm"
)

type LimitRepositoryImpl struct {
	db *gorm.DB
}

func NewLimitRepository(db *gorm.DB) LimitRepository {
	return &LimitRepositoryImpl{db: db}
}

func (r *LimitRepositoryImpl) FindAll() ([]models.Limit, error) {
	var limits []models.Limit
	if err := r.db.Find(&limits).Error; err != nil {
		return nil, err
	}
	return limits, nil
	
}

func (r *LimitRepositoryImpl) FindById(id int) (models.Limit, error) {
	var limit models.Limit
	if err := r.db.First(&limit, id).Error; err != nil {
		return models.Limit{}, err
	}
	return limit, nil
	
}

func (r *LimitRepositoryImpl) Save(limit models.Limit) (error) {
	err := r.db.Save(&limit).Error
	return err
}

func (r *LimitRepositoryImpl) Update(limit models.Limit) (models.Limit, error) {
	err := r.db.Save(&limit).Error
	return limit, err
	
}

func (r *LimitRepositoryImpl) Delete(limit models.Limit) (error) {
	err := r.db.Delete(&limit).Error
	return err
}

func (r *LimitRepositoryImpl) FindByKonsumenTenor(konsumenID int, tenor uint8) (models.Limit, error) {
	var limits models.Limit
	if err := r.db.Where("konsumen_id = ?" + " AND tenor = ?", konsumenID, tenor).First(&limits).Error; err != nil {
		return models.Limit{}, err
	}
	return limits, nil
	
}