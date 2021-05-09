package entity

type PilotEntity struct {
	ID             uint               `gorm:"primaryKey;autoIncrement" json:"id"`
	Name           string             `json:"name"`
	OrganizationID uint               `json:"organizationId"`
	TimePart       BaseTimeEntity     `gorm:"embedded" json:"-"`
	Organization   OrganizationEntity `json:"organization"`
	Gundam         []GundamEntity     `gorm:"many2many:gundam_pilot"`
}

func (PilotEntity) TableName() string {
	return "pilot"
}
