package dto

type CreateLimitRequest struct {
	KonsumenID  int    `json:"konsumen_id" binding:"required"`
	Tenor       uint8  `json:"tenor" binding:"required,oneof=1 2 3 4"` // hanya boleh 1–4
	LimitAmount string `json:"limit_amount" binding:"required"`       
}

type UpdateLimitRequest struct {
	Tenor       uint8  `json:"tenor"`
	LimitAmount string `json:"limit_amount"`
}
