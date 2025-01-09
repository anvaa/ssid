package users

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"
)

var UsrDB *gorm.DB

func CnnUserDB(dbpath string) {
	var err error
	UsrDB, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	UsrDB.Exec("PRAGMA foreign_keys = ON;")
	// UsrDB.Exec("PRAGMA journal_mode=WAL;")

	SyncUserDB(UsrDB)
	log.Println("UsrDB connected and synced")
}

func CloseUserDB() {
	sqlDB, err := UsrDB.DB()
	if err != nil {
		log.Fatal("failed to close database")
	}
	sqlDB.Close()
	log.Println("UserDB closed")
}