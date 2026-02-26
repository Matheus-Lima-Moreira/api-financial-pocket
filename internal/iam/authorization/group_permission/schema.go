package group_permission

import (
	"time"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/authorization/action"
)

type GroupPermissionSchema struct {
	ID        uint                `gorm:"primaryKey"`
	Name      string              `gorm:"not null;size:255"`
	Type      GroupPermissionType `gorm:"not null;size:255"`
	CreatedAt time.Time           `gorm:"autoCreateTime"`
	UpdatedAt time.Time           `gorm:"autoUpdateTime"`

	GroupPermissionActions []GroupPermissionActionSchema `gorm:"foreignKey:GroupPermissionID"`
}

func (GroupPermissionSchema) TableName() string {
	return "group_permissions"
}

type GroupPermissionActionSchema struct {
	ID                uint      `gorm:"primaryKey"`
	GroupPermissionID uint      `gorm:"not null"`
	ActionID          uint      `gorm:"not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`

	GroupPermission GroupPermissionSchema `gorm:"foreignKey:GroupPermissionID"`
	Action          action.ActionSchema   `gorm:"foreignKey:ActionID"`
}

func (GroupPermissionActionSchema) TableName() string {
	return "group_permission_actions"
}
