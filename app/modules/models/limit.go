package models

import "time"

type Limit struct {
	ID          int       `gorm:"primaryKey;autoIncrement" json:"id"`
	KonsumenID  int       `gorm:"not null;index" json:"konsumen_id"`
	Tenor       uint8     `gorm:"type:tinyint;not null;check:tenor IN (1,2,3,4)" json:"tenor"`
	LimitAmount int64     `gorm:"type:bigint;not null" json:"limit_amount"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	Konsumen *Konsumen `gorm:"foreignKey:KonsumenID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"konsumen,omitempty"`
}
