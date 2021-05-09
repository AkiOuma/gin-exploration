package entity

type OrganizationEntity struct {
	ID       uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string         `json:"name"`
	Alias    string         `json:"alias"`
	TimePart BaseTimeEntity `gorm:"embedded" json:"-"`
	Pilot    []PilotEntity  `gorm:"foreignKey:OrganizationID" json:"pilot"`
}

func (OrganizationEntity) TableName() string {
	return "organization"
}
