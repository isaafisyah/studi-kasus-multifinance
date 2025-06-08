package dto

import "mime/multipart"

type CreateKonsumenRequest struct {
	NIK          string `form:"nik" binding:"required" minlength:"16" maxlength:"16"`
	Fullname     string `form:"fullname" binding:"required" minlength:"3"`
	LegalName    string `form:"legal_name" binding:"required" minlength:"3"`
	TempatLahir  string `form:"tempat_lahir" binding:"required" minlength:"3"`
	TanggalLahir string `form:"tanggal_lahir" binding:"required"` 
	Gaji         string `form:"gaji" binding:"required"`    
	FotoKtp      *multipart.FileHeader `form:"foto_ktp"`
	FotoSelfie   *multipart.FileHeader `form:"foto_selfie"`      
}

type UpdateKonsumenRequest struct {
	Fullname     *string `form:"fullname" minlength:"3"`
	LegalName    *string `form:"legal_name" minlength:"3"`
	TempatLahir  *string `form:"tempat_lahir" minlength:"3"`
	TanggalLahir *string `form:"tanggal_lahir"`
	Gaji         *string `form:"gaji"`
}
