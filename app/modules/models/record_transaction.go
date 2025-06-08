package models

import "time"

type RecordTransaction struct {
	ID            int       `gorm:"primaryKey;autoIncrement" json:"id"`
	KonsumenID    int       `gorm:"not null;index" json:"konsumen_id"`
	NomorKontrak  string    `gorm:"type:varchar(50);not null" json:"nomor_kontrak"`
	OTR           string     `gorm:"type:bigint;not null" json:"otr"`
	AdminFee      int64     `gorm:"type:bigint;not null" json:"admin_fee"`
	JumlahCicilan int64     `gorm:"type:bigint;not null" json:"jumlah_cicilan"`
	JumlahBunga   int64     `gorm:"type:bigint;not null" json:"jumlah_bunga"`
	NamaAset      string    `gorm:"type:varchar(255);not null" json:"nama_aset"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Konsumen *Konsumen `gorm:"foreignKey:KonsumenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"konsumen,omitempty"`
}
