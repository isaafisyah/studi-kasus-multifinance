package models

import "time"

type Konsumen struct{
	ID         int       `gorm:"primaryKey;autoIncrement"`
    NIK        string    `gorm:"type:varchar(16);unique;not null" json:"nik"`
    Fullname   string    `gorm:"type:varchar(255);not null" json:"fullname"`
    LegalName  string    `gorm:"type:varchar(255);not null" json:"legal_name"`
    TempatLahir string   `gorm:"type:varchar(255);not null" json:"tempat_lahir"`
    TanggalLahir time.Time   `gorm:"type:varchar(255);not null" json:"tanggal_lahir"`
    Gaji       int64    `gorm:"type:bigint;not null" json:"gaji"`
	FotoKtp    string    `gorm:"type:varchar(255);not null" json:"foto_ktp"`
	FotoSelfie string    `gorm:"type:varchar(255);not null" json:"foto_selfie"`
    CreatedAt  time.Time `gorm:"autoCreateTime"`
    UpdatedAt  time.Time `gorm:"autoUpdateTime"`

    // Limits     []Limit   `gorm:"foreignKey:KonsumenID"`
	// RecordTransactions []RecordTransaction `gorm:"foreignKey:KonsumenID"`
}