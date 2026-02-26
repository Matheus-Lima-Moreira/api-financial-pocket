package group_permission

import (
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&GroupPermissionSchema{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&GroupPermissionActionSchema{}); err != nil {
		return err
	}

	return nil
}
