package action

import (
	"time"
)

type ActionSchema struct {
	ID        uint      `gorm:"primaryKey"`
	Resource  string    `gorm:"not null;size:255"`
	Action    string    `gorm:"not null;size:255"`
	Label     string    `gorm:"not null;size:255"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (ActionSchema) TableName() string {
	return "actions"
}
