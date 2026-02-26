package user

import (
	"time"

	group_permission "github.com/Matheus-Lima-Moreira/financial-pocket/internal/iam/authorization/group_permission"
)

type UserSchema struct {
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"not null;size:255"`
	Email         string    `gorm:"uniqueIndex;not null;size:255"`
	Password      string    `gorm:"not null;size:255"`
	EmailVerified bool      `gorm:"not null;default:false"`
	Avatar        string    `gorm:"not null;size:255"`
	CreatedAt     time.Time `gorm:"autoCreateTime"`
	UpdatedAt     time.Time `gorm:"autoUpdateTime"`

	UserGroupPermissions []UserGroupPermissionSchema `gorm:"foreignKey:UserID"`
}

func (UserSchema) TableName() string {
	return "users"
}

type UserGroupPermissionSchema struct {
	ID                uint      `gorm:"primaryKey"`
	UserID            uint      `gorm:"not null"`
	GroupPermissionID uint      `gorm:"not null"`
	CreatedAt         time.Time `gorm:"autoCreateTime"`
	UpdatedAt         time.Time `gorm:"autoUpdateTime"`

	User            UserSchema                             `gorm:"foreignKey:UserID"`
	GroupPermission group_permission.GroupPermissionSchema `gorm:"foreignKey:GroupPermissionID"`
}

func (UserGroupPermissionSchema) TableName() string {
	return "user_group_permissions"
}
