package repositories

import "github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"

type RecordTransactionRepository interface {
	FindAll() ([]models.RecordTransaction, error)
	FindById(id int) (models.RecordTransaction, error)
	Save(recordTransaction models.RecordTransaction) (error)
	FindByKonsumenID(id int) ([]models.RecordTransaction, error)
}