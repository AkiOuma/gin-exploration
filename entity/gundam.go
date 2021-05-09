package entity

type GundamEntity struct {
	ID       uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Code     string         `json:"code"`
	Name     string         `json:"name"`
	TimePart BaseTimeEntity `gorm:"embedded" json:"-"`
	Pilot    []PilotEntity  `gorm:"many2many:gundam_pilot" json:"pilot"`
}

func (GundamEntity) TableName() string {
	return "gundam"
}
