package services

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/isaafisyah/studi-kasus-multifinance/app/log"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/repositories"
)

type RecordTransactionServiceImpl struct {
	recordTransactionRepository repositories.RecordTransactionRepository
	limitRepository repositories.LimitRepository
}

func NewRecordTransactionService(recordTransactionRepository repositories.RecordTransactionRepository, limitRepository repositories.LimitRepository) RecordTransactionService {
	return &RecordTransactionServiceImpl{recordTransactionRepository : recordTransactionRepository, limitRepository: limitRepository}
}

func (rt *RecordTransactionServiceImpl) FindAll() ([]models.RecordTransaction, error) {
	return rt.recordTransactionRepository.FindAll()
	
}

func (rt *RecordTransactionServiceImpl) FindById(idStr string) (models.RecordTransaction, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return models.RecordTransaction{}, err
	}
	return rt.recordTransactionRepository.FindById(id)
	
}

func (rt *RecordTransactionServiceImpl) Create(req dto.CreateRecordTransactionRequest) (*models.RecordTransaction,error) {
	log.GetLogger("TransactionService").Info("Service Create Transaction Start")
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	
	var recordTransaction models.RecordTransaction
	var errResult error
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		mu.Lock()
		defer mu.Unlock()
	
		adminFee,_ := strconv.ParseInt(req.AdminFee, 10, 64)
		jumlahBunga,_ := strconv.ParseInt(req.JumlahBunga, 10, 64)
		cicilan,_ := strconv.ParseInt(req.JumlahCicilan, 10, 64)
		jumlahCicilan :=cicilan+ adminFee + jumlahBunga
	
		var existingCount uint8
		limit := uint8(4)
	
		transaction, _ := rt.recordTransactionRepository.FindByKonsumenID(req.KonsumenID)
		log.GetLogger("TransactionService").Debug(transaction)
		if len(transaction) == 0 {
			existingCount = 1
		} else {
			existingCount = uint8(len(transaction))
		}
	
		limits, err := rt.limitRepository.FindByKonsumenTenor(req.KonsumenID, existingCount)
		log.GetLogger("TransactionService").Debug(limits)
		if err != nil {
			errResult = fmt.Errorf("cicilan tidak ada")
			log.GetLogger("TransactionService").Error(errResult)
			return
		}
	
		if jumlahCicilan > limits.LimitAmount {
			errResult = fmt.Errorf("jumlah cicilan melebihi batas limit yang ditentukan")
			log.GetLogger("TransactionService").Error(errResult)
			return
		}

		if jumlahCicilan < limits.LimitAmount {
			errResult = fmt.Errorf("jumlah cicilan kurang dari batas limit yang ditentukan")
			log.GetLogger("TransactionService").Error(errResult)
			return
		}
	
		if existingCount >= limit {
			errResult = fmt.Errorf("jumlah cicilan sudah mencapai batas maksimal (%d)", limit)
			log.GetLogger("TransactionService").Error(errResult)
			return
		}
	
		recordTransaction = models.RecordTransaction{
			KonsumenID:    req.KonsumenID,
			NomorKontrak:  req.NomorKontrak,
			OTR:           req.OTR,
			AdminFee:      adminFee,
			JumlahCicilan: jumlahCicilan,
			JumlahBunga:   jumlahBunga,
			NamaAset:      req.NamaAset,
		}
	}()
	
	wg.Wait()
	
	if errResult != nil {
		return nil, errResult
	}
	log.GetLogger("TransactionService").Debug(recordTransaction)
	log.GetLogger("TransactionService").Info("Service Create Transaction End")
	
	return &recordTransaction, nil
	
}
	