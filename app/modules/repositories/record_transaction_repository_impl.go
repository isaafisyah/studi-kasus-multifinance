package repositories

import (
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"gorm.io/gorm"
)

type RecordTransactionRepositoryImpl struct {
	db *gorm.DB
}

func NewRecordTransactionRepository(db *gorm.DB) RecordTransactionRepository {
	return &RecordTransactionRepositoryImpl{db: db}
}

func (r *RecordTransactionRepositoryImpl) FindAll() ([]models.RecordTransaction, error) {
	var transactions []models.RecordTransaction
	if err := r.db.Find(&transactions).Error; err != nil {
		return nil, err
	}
	return transactions, nil
	
}

func (r *RecordTransactionRepositoryImpl) FindById(id int) (models.RecordTransaction, error) {
	var transaction models.RecordTransaction
	if err := r.db.First(&transaction, id).Error; err != nil {
		return models.RecordTransaction{}, err
	}
	return transaction, nil
	
}

func (r *RecordTransactionRepositoryImpl) Save(transaction models.RecordTransaction) (error) {
	err := r.db.Save(&transaction).Error
	return err
}

func (r *RecordTransactionRepositoryImpl) FindByKonsumenID(id int) ([]models.RecordTransaction, error) {
	var recordTransaction []models.RecordTransaction
	if err := r.db.Where("konsumen_id = ?", id).Find(&recordTransaction).Error; err != nil {
		return nil, err
	}
	
	return recordTransaction, nil
}