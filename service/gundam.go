package service

import (
	"gin-demo01/dto"
	"gin-demo01/entity"
	"gin-demo01/utils"

	"gorm.io/gorm"
)

func CreateGundam(gundam entity.GundamEntity) {
	entity.DB.Select("Code", "Name").Create(&gundam)
}

func BatchCreateGundam(gundam []entity.GundamEntity) {
	entity.DB.Omit("DeletedAt").Create(&gundam)
}

func GetGundamList(query dto.QueryGundamDTO) (gundam []entity.GundamEntity) {
	db := entity.DB.Preload("Pilot", func(db *gorm.DB) *gorm.DB {
		return db.Preload("Organization")
	})
	utils.LikeOrEqualQuery(db, query)
	utils.Pagination(db, query.Page, query.PageSize)
	db.Find(&gundam)
	return gundam
}

func UpdateGundam(gundam entity.GundamEntity) {
	entity.DB.Model(&gundam).Updates(&gundam)
}

func DeleteGundam(id uint) {
	entity.DB.Delete(&entity.GundamEntity{ID: id})
}
