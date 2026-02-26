package user

import "gorm.io/gorm"

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&UserSchema{}); err != nil {
		return err
	}

	if err := db.AutoMigrate(&UserGroupPermissionSchema{}); err != nil {
		return err
	}

	return nil
}
