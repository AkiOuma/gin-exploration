package service

import (
	"gin-demo01/entity"
)

func CreateOrganization(organization entity.OrganizationEntity) {
	entity.DB.Select("Name", "Alias").Create(&organization)
}

func GetOrganizationList() (organization []entity.OrganizationEntity) {
	entity.DB.Preload("Pilot").Find(&organization)
	return organization
}
