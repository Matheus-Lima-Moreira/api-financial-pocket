package organizations

import "time"

type OrganizationSchema struct {
	ID        string    `gorm:"primaryKey;not null"`
	Name      string    `gorm:"not null;size:255"`
	Cellphone string    `gorm:"not null;size:255"`
	Logo      string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (OrganizationSchema) TableName() string {
	return "organizations"
}
