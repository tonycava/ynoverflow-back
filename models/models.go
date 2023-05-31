package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

func GenerateISOString() string {
	return time.Now().Format("2006-01-02T15:04:05.999Z07:00")
}

type Base struct {
	ID        string `gorm:"primarykey" json:"id"`
	UUID      string `gorm:"primaryKey;autoIncrement:false" json:"uuid"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.UUID = uuid.New().String()

	t := GenerateISOString()
	base.CreatedAt, base.UpdatedAt = t, t

	return nil
}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	// update timestamps
	base.UpdatedAt = GenerateISOString()
	return nil
}
