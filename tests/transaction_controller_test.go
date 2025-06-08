package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateRecordTransaction_Success(t *testing.T) {
    gin.SetMode(gin.TestMode)

    mockService := new(MockRecordTransactionService)
	mockRepo := new(MockRecordTransactionRepository)
	controller := controllers.NewRecordTransactionController(mockService, mockRepo)

    dummyReq := dto.CreateRecordTransactionRequest{
        KonsumenID:    1,
        NomorKontrak:  "kontrak123",
        OTR:           "otr123",
        AdminFee:      "10000",
        JumlahCicilan: "1050000",
        JumlahBunga:   "50000",
        NamaAset:      "Mobil",
    }

    dummyRes := &models.RecordTransaction{
        ID:           1,
        KonsumenID:   1,
        NomorKontrak: "kontrak123",
		OTR:          "otr123",
		AdminFee:     10000,
		JumlahCicilan:1050000,
		JumlahBunga:  50000,
        NamaAset:     "Mobil",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
    }

    mockService.On("Create", dummyReq).Return(dummyRes, nil)

	mockRepo.On("Save", *dummyRes).Return(nil)

    // JSON encode request body
    reqBody, _ := json.Marshal(dummyReq)
	router := gin.Default()
	router.POST("/transaction", controller.Create)
	req := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
    mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}

func TestCreateRecordTransaction_Failed(t *testing.T) {
    gin.SetMode(gin.TestMode)

    mockService := new(MockRecordTransactionService)
	mockRepo := new(MockRecordTransactionRepository)
	controller := controllers.NewRecordTransactionController(mockService, mockRepo)

    dummyReq := dto.CreateRecordTransactionRequest{
        KonsumenID:    1,
        NomorKontrak:  "kontrak123",
        OTR:           "otr123",
        AdminFee:      "10000",
        JumlahCicilan: "1050000",
        JumlahBunga:   "50000",
        NamaAset:      "",
    }

    dummyRes := &models.RecordTransaction{
        ID:           1,
        KonsumenID:   1,
        NomorKontrak: "kontrak123",
		OTR:          "otr123",
		AdminFee:     10000,
		JumlahCicilan:1050000,
		JumlahBunga:  50000,
        NamaAset:     "Mobil",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
    }

    mockService.On("Create", dummyReq).Return(dummyRes, nil)

	mockRepo.On("Save", *dummyRes).Return(nil)

    // JSON encode request body
    reqBody, _ := json.Marshal(dummyReq)
	router := gin.Default()
	router.POST("/transaction", controller.Create)
	req := httptest.NewRequest(http.MethodPost, "/transaction", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusBadRequest, rr.Code)
}
