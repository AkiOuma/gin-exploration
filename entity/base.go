package entity

import (
	"time"

	"gorm.io/gorm"
)

type BaseTimeEntity struct {
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

type Tabler interface {
	TableName()
}
