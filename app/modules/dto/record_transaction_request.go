package dto

type CreateRecordTransactionRequest struct {
	KonsumenID    int    `json:"konsumen_id" binding:"required"`
	NomorKontrak  string `json:"nomor_kontrak" binding:"required"`
	OTR           string `json:"otr" binding:"required"`
	AdminFee      string `json:"admin_fee" binding:"required"`
	JumlahCicilan string `json:"jumlah_cicilan" binding:"required"`
	JumlahBunga   string `json:"jumlah_bunga" binding:"required"`
	NamaAset      string `json:"nama_aset" binding:"required"`
}