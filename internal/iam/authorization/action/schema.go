package action

type ActionSchema struct {
	ID          string `gorm:"primaryKey;not null"`
	Resource    string `gorm:"not null;size:255"`
	Action      string `gorm:"not null;size:255"`
	Label       string `gorm:"not null;size:255"`
	Description string `gorm:"not null;size:255"`
}

func (ActionSchema) TableName() string {
	return "actions"
}
