package app_db

import (
	"gorm.io/gorm"

	"app/app_models"
	
)

func SyncAppDB(db *gorm.DB) {
	db.AutoMigrate(
		&app_models.Items{},
		&app_models.LocNames{},
		&app_models.ManNames{},
		&app_models.TypNames{},
		&app_models.StaNames{},
		&app_models.Status_History{},
		
	)
}