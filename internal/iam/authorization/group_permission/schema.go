package group_permission

import (
	"time"

	"github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/authorization/action"
	organization "github.com/Matheus-Lima-Moreira/financial-pocket/internal/organizations"
)

type GroupPermissionSchema struct {
	ID             string              `gorm:"primaryKey"`
	Name           string              `gorm:"not null;size:255"`
	Type           GroupPermissionType `gorm:"not null;size:255"`
	OrganizationID string
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	GroupPermissionActions []GroupPermissionActionSchema   `gorm:"foreignKey:GroupPermissionID"`
	Organization           organization.OrganizationSchema `gorm:"foreignKey:OrganizationID"`
}

func (GroupPermissionSchema) TableName() string {
	return "group_permissions"
}

type GroupPermissionActionSchema struct {
	GroupPermissionID string `gorm:"not null"`
	ActionID          string `gorm:"not null"`

	GroupPermission GroupPermissionSchema `gorm:"foreignKey:GroupPermissionID"`
	Action          action.ActionSchema   `gorm:"foreignKey:ActionID"`
}

func (GroupPermissionActionSchema) TableName() string {
	return "group_permission_actions"
}
