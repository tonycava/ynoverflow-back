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
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) error {
	base.ID = uuid.New().String()

	t := GenerateISOString()
	base.CreatedAt = t
	base.UpdatedAt = t
	return nil
}

func (base *Base) AfterUpdate(tx *gorm.DB) error {
	base.UpdatedAt = GenerateISOString()
	return nil
}
