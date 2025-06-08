package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/dto"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateLimitController(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockService := new(MockLimitService)
	mockRepo := new(MockLimitRepository)
	controller := controllers.NewLimitController(mockService, mockRepo)

	dummyReq := dto.CreateLimitRequest{
		KonsumenID:    1,
		Tenor:         1,
		LimitAmount: "1000000",
	}

	dummyRes := &models.Limit{
		ID:          1,
		KonsumenID:  1,
		Tenor:       1,
		LimitAmount: 1000000,
	}

	mockService.On("Create", dummyReq).Return(dummyRes, nil)

	mockRepo.On("Save", *dummyRes).Return(nil)

	reqBody, _ := json.Marshal(dummyReq)
	router := gin.Default()
	router.POST("/limit", controller.Create)
	req := httptest.NewRequest(http.MethodPost, "/limit", bytes.NewBuffer(reqBody))
    req.Header.Set("Content-Type", "application/json")

    rr := httptest.NewRecorder()
    router.ServeHTTP(rr, req)

    assert.Equal(t, http.StatusCreated, rr.Code)
    mockService.AssertExpectations(t)
	mockRepo.AssertExpectations(t)
}