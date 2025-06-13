package tests

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/isaafisyah/studi-kasus-multifinance/app/modules/controllers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateKonsumen_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockKonsumenService)
	controller := controllers.NewKonsumenController(mockService)

	// Dummy file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("nik", "1234567890123456")
	_ = writer.WriteField("fullname", "Budi")
	_ = writer.WriteField("legal_name", "Budi")
	_ = writer.WriteField("tempat_lahir", "Jakarta")
	_ = writer.WriteField("tanggal_lahir", "2000-01-01")
	_ = writer.WriteField("gaji", "10000000")

	part, _ := writer.CreateFormFile("foto_ktp", "ktp.jpg")
	part.Write([]byte("dummy image content"))

	part2, _ := writer.CreateFormFile("foto_selfie", "selfie.jpg")
	part2.Write([]byte("dummy selfie content"))

	writer.Close()

	// dummyRes := &models.Konsumen{
	// 	ID:           1,
	// 	NIK:          "1234567890123456",
	// 	Fullname:     "Budi",
	// 	LegalName:    "Budi",
	// 	TempatLahir:  "Jakarta",
	// 	TanggalLahir: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Gaji:         10000000,
	// }

	mockService.On("Create", mock.AnythingOfType("dto.CreateKonsumenRequest"), mock.AnythingOfType("*gin.Context")).Return(nil)

	router := gin.Default()
	router.POST("/konsumen", controller.Create)
	req := httptest.NewRequest(http.MethodPost, "/konsumen", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	mockService.AssertExpectations(t)
}

func TestCreateKonsumen_Failed(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockKonsumenService)
	controller := controllers.NewKonsumenController(mockService)

	// Dummy file
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("nik", "1234567890123456")
	_ = writer.WriteField("fullname", "Budi")
	_ = writer.WriteField("legal_name", "Budi")
	_ = writer.WriteField("tempat_lahir", "Jakarta")
	_ = writer.WriteField("tanggal_lahir", "2000-01-01")
	_ = writer.WriteField("gaji", "")

	part, _ := writer.CreateFormFile("foto_ktp", "ktp.jpg")
	part.Write([]byte("dummy image content"))

	part2, _ := writer.CreateFormFile("foto_selfie", "selfie.jpg")
	part2.Write([]byte("dummy selfie content"))

	writer.Close()

	// dummyRes := &models.Konsumen{
	// 	ID:           1,
	// 	NIK:          "1234567890123456",
	// 	Fullname:     "Budi",
	// 	LegalName:    "Budi",
	// 	TempatLahir:  "Jakarta",
	// 	TanggalLahir: time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	Gaji:         10000000,
	// }

	mockService.On("Create", mock.AnythingOfType("dto.CreateKonsumenRequest"), mock.AnythingOfType("*gin.Context")).Return(nil)

	router := gin.Default()
	router.POST("/konsumen", controller.Create)
	req := httptest.NewRequest(http.MethodPost, "/konsumen", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}
