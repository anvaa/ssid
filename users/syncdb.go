package users

import (
	"gorm.io/gorm"
)

func SyncUserDB(db *gorm.DB) {
	db.AutoMigrate(
		&Users{},
		&Links{},
	)
}