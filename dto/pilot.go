package dto

type PilotDTO struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	OrganizationID uint   `json:"organizationId"`
}

type QueryPilotDTO struct {
	Name           string `form:"name" mode:"like" field:"name"`
	OrganizationID int    `form:"organizationId" mode:"equal" field:"organization_id"`
	Page           int    `form:"page" json:"page" binding:"required" `
	PageSize       int    `form:"pageSize" json:"pageSize" binding:"required"`
}
