package group_permission

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	return db.
		Where("type = ?", GroupPermissionSystem).
		FirstOrCreate(&GroupPermissionSchema{
			Name: "System",
			Type: GroupPermissionSystem,
		}).Error
}
