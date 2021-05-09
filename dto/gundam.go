package dto

import "gorm.io/gorm"

type GetGundamDTO struct {
	ID        uint           `json:"id"`
	Code      string         `json:"code"`
	Name      string         `json:"name"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

type QueryGundamDTO struct {
	Code     string `form:"code" mode:"like" field:"code"`
	Name     string `form:"name" mode:"like" field:"name"`
	Page     int    `form:"page" json:"page" binding:"required" `
	PageSize int    `form:"pageSize" json:"pageSize" binding:"required"`
}
