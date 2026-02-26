package action

import (
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	return db.
		Where("resource = ?", "users").
		Where("action = ?", "create").
		FirstOrCreate(&ActionSchema{
			Resource: "users",
			Action:   "create",
			Label:    "Create User",
		}).Error
}
