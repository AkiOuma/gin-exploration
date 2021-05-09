package service

import (
	"gin-demo01/dto"
	"gin-demo01/entity"
	"gin-demo01/utils"
)

func GetPilotList(query dto.QueryPilotDTO, pagination dto.PaginationDTO) []entity.PilotEntity {
	var pilots []entity.PilotEntity
	db := entity.DB.Preload("Organization").Preload("Gundam")
	utils.LikeOrEqualQuery(db, query)
	utils.Pagination(db, pagination.Page, pagination.PageSize)
	db.Find(&pilots)
	return pilots
}

func CreatePilot(pilot entity.PilotEntity) {
	entity.DB.Select("Name", "OrganizationID").Create(&pilot)
}
