package repositories

import "github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"

type LimitRepository interface{
	FindAll() ([]models.Limit, error)
	FindById(id int) (models.Limit, error)
	Save(limit models.Limit) (error)
	Update(limit models.Limit) (models.Limit, error)
	Delete(limit models.Limit) (error)
	FindByKonsumenTenor(konsumenID int, tenor uint8) (models.Limit, error)
}