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

	if err := userDBInit(UsrDB); err != nil {
		log.Fatal(err)
	}

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

func userDBInit(db *gorm.DB) error {
	if err := syncUserDB(db); err != nil {
		return err
	}

	// check if admin user exists
	var count int64
	if err := db.Model(&Users{}).Where("role = ?", "admin").Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		if err := insertDefaultData(db); err != nil {
			return err
		}
	}

	return nil
}

func insertDefaultData(db *gorm.DB) error {
	adminUser := Users{
		Id:       1212090603,
		Email:    "admin@ssid.loc",
		Password: "$2a$10$10ZlTiAVW7EkKMp4559RPuv91.O9tLO7cx6azy72W8AuCBDST8.de",
		Role:     "admin",
		IsAuth:   true,
		AccessTime:  24600,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	link := Links{
		Id:     1212090602,
		Url:    "/v/newusers",
		UserId: adminUser.Id,
	}

	if err := db.Create(&link).Error; err != nil {
		return err
	}

	return nil
}

func syncUserDB(db *gorm.DB) error {
	SyncUserDB(db)
	return nil
}