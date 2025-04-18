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
		if err := InsertDefaultData(db); err != nil {
			return err
		}
	}

	return nil
}

func syncUserDB(db *gorm.DB) error {
	SyncUserDB(db)
	return nil
}

func InsertDefaultData(db *gorm.DB) error {
	if err := insertAdminData(db); err != nil {
		return err
	}

	if err := insertUserData(db); err != nil {
		return err
	}

	return nil
}

func insertAdminData(db *gorm.DB) error {
	adminUser := Users{
		Id:       1212090603,
		Email:    "admin@ssid.loc",
		// ssidadmin25
		Password: "$2a$10$FZoSUNhpWs9L1MXS3GwTA.1FF2K5ICaTzJgKKmda513hTNRYYrV4m",
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
		UserId: 1212090603,
	}

	if err := db.Create(&link).Error; err != nil {
		return err
	}

	return nil
}

func insertUserData(db *gorm.DB) error {
	adminUser := Users{
		Id:       1212090602,
		Email:    "user@ssid.loc",
		// ssiduser25
		Password: "$2a$10$UGMEocTQjWkNArdx2qsyiOA6yWS69Qq7o78iMrhJKmz6vpkeKiwzy",
		Role:     "user",
		IsAuth:   true,
		AccessTime:  24600,
	}

	if err := db.Create(&adminUser).Error; err != nil {
		return err
	}

	link := Links{
		Id:     1212090601,
		Url:    "/app",
		UserId: 1212090602,
	}

	if err := db.Create(&link).Error; err != nil {
		return err
	}

	return nil
}

