package tests

import (
	"fmt"
	"testing"
	"time"

	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/services"
	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction_LimitValid(t *testing.T) {
	mockLimitRepo := new(MockLimitRepository)
	mockRTRepo := new(MockRecordTransactionRepository)

	service := services.NewRecordTransactionService(mockRTRepo, mockLimitRepo)

	req := dto.CreateRecordTransactionRequest{
		KonsumenID:     1,
		NomorKontrak:   "ctr001",
		OTR:            "2348032",
		AdminFee:       "100000",
		JumlahCicilan:  "9000000",
		JumlahBunga:    "100000",
		NamaAset:       "Mobil",
	}

	mockRTRepo.On("FindByKonsumenID", 1).Return([]models.RecordTransaction{}, nil)

	mockLimitRepo.On("FindByKonsumenTenor", 1, uint8(1)).Return(models.Limit{
		LimitAmount: 20000000,
	}, nil)

	result, err := service.Create(req)
	assert.NoError(t, err)
	assert.Equal(t, req.KonsumenID, result.KonsumenID)
	assert.Equal(t, req.NamaAset, result.NamaAset)
}

func TestRecordTransactionService_Create_LimitNotValid(t *testing.T) {
    mockLimitRepo := new(MockLimitRepository)
	mockRTRepo := new(MockRecordTransactionRepository)

    service := services.NewRecordTransactionService(mockRTRepo, mockLimitRepo)

    req := dto.CreateRecordTransactionRequest{
        KonsumenID:   1,
        NomorKontrak: "kontrak123",
        OTR:          "otr123",
        AdminFee:     "10000",
        JumlahBunga:  "50000",
        JumlahCicilan:"200000",
        NamaAset:     "Mobil",
    }

    mockRTRepo.On("FindByKonsumenID", req.KonsumenID).
        Return([]models.RecordTransaction{ {} }, nil)

    mockLimitRepo.On("FindByKonsumenTenor", req.KonsumenID, uint8(1)).
        Return(models.Limit{}, fmt.Errorf("db error"))

    record, err := service.Create(req)

    assert.Nil(t, record)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), "cicilan tidak ada")

    mockRTRepo.AssertExpectations(t)
    mockLimitRepo.AssertExpectations(t)
}

func TestRecordTransactionService_Create_OverLimitAmount(t *testing.T) {
	mockLimitRepo := new(MockLimitRepository)
	mockRTRepo := new(MockRecordTransactionRepository)

	service := services.NewRecordTransactionService(mockRTRepo, mockLimitRepo)

	req := dto.CreateRecordTransactionRequest{
		KonsumenID:    1,
		NomorKontrak:  "TRX-999",
		OTR:           "ot234dsaf",
		AdminFee:      "10000",
		JumlahCicilan: "1050000", 
		JumlahBunga:   "50000",
		NamaAset:      "Motor Mewah",
	}

	expectedCicilan := int64(10000 + 1050000 + 50000) 

	mockRTRepo.On("FindByKonsumenID", req.KonsumenID).Return([]models.RecordTransaction{
		{ID: 1},
	}, nil)

	mockLimitRepo.On("FindByKonsumenTenor", req.KonsumenID, uint8(1)).Return(models.Limit{
		ID:          1,
		KonsumenID:  1,
		Tenor:       1,
		LimitAmount: expectedCicilan - 5000, // batas limit 1.15jt
	}, nil)

	// Run
	result, err := service.Create(req)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "jumlah cicilan melebihi batas limit yang ditentukan")
}

func TestRecordTransactionService_Create_LessLimitAmount(t *testing.T) {
	mockLimitRepo := new(MockLimitRepository)
	mockRTRepo := new(MockRecordTransactionRepository)

	service := services.NewRecordTransactionService(mockRTRepo, mockLimitRepo)

	req := dto.CreateRecordTransactionRequest{
		KonsumenID:    1,
		NomorKontrak:  "TRX-999",
		OTR:           "ot234dsaf",
		AdminFee:      "10000",
		JumlahCicilan: "1050000", 
		JumlahBunga:   "50000",
		NamaAset:      "Motor Mewah",
	}

	expectedCicilan := int64(10000 + 1050000 + 50000) 

	mockRTRepo.On("FindByKonsumenID", req.KonsumenID).Return([]models.RecordTransaction{
		{ID: 1},
	}, nil)

	mockLimitRepo.On("FindByKonsumenTenor", req.KonsumenID, uint8(1)).Return(models.Limit{
		ID:          1,
		KonsumenID:  1,
		Tenor:       1,
		LimitAmount: expectedCicilan + 5000, 
	}, nil)

	// Run
	result, err := service.Create(req)

	// Assert
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "jumlah cicilan kurang dari batas limit yang ditentukan")
}

func TestRecordTransactionService_Create_MaximumLimit(t *testing.T) {
    mockLimitRepo := new(MockLimitRepository)
	mockRTRepo := new(MockRecordTransactionRepository)

    service := services.NewRecordTransactionService(mockRTRepo, mockLimitRepo)

    req := dto.CreateRecordTransactionRequest{
        KonsumenID:   1,
        NomorKontrak: "kontrak123",
        OTR:          "otr123",
        AdminFee:     "10000",
        JumlahBunga:  "1050000",
        JumlahCicilan:"50000",
        NamaAset:     "Mobil",
    }

	dummyRecordTransactions := []models.RecordTransaction{
		{
			ID:            1,
			KonsumenID:    1,
			NomorKontrak:  "kontrak1",
			OTR:           "otr123",
			AdminFee:      10000,
			JumlahCicilan: 1050000,
			JumlahBunga:   50000,
			NamaAset:      "Mobil A",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            2,
			KonsumenID:    1,
			NomorKontrak:  "kontrak2",
			OTR:           "otr124",
			AdminFee:      10000,
			JumlahCicilan: 1050000,
			JumlahBunga:   50000,
			NamaAset:      "Mobil B",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            3,
			KonsumenID:    1,
			NomorKontrak:  "kontrak3",
			OTR:           "otr125",
			AdminFee:      10000,
			JumlahCicilan: 1050000,
			JumlahBunga:   50000,
			NamaAset:      "Mobil C",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
		{
			ID:            4,
			KonsumenID:    1,
			NomorKontrak:  "kontrak4",
			OTR:           "otr126",
			AdminFee:      10000,
			JumlahCicilan: 1050000,
			JumlahBunga:   50000,
			NamaAset:      "Mobil D",
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
		},
	}

	expectedCicilan := int64(10000 + 1050000 + 50000) 
	
	// Mock return dari repository
	mockRTRepo.On("FindByKonsumenID", req.KonsumenID).
		Return(dummyRecordTransactions, nil)
	

		mockLimitRepo.On("FindByKonsumenTenor", req.KonsumenID, uint8(4)).Return(models.Limit{
			ID:          1,
			KonsumenID:  1,
			Tenor:       4,
			LimitAmount: expectedCicilan, 
		}, nil)

    record, err := service.Create(req)

    assert.Nil(t, record)
    assert.Error(t, err)
    assert.Contains(t, err.Error(), fmt.Sprintf("jumlah cicilan sudah mencapai batas maksimal (%d)", 4))

    mockRTRepo.AssertExpectations(t)
    mockLimitRepo.AssertExpectations(t)
}


