package app_db

import (
	"log"
	"srv/filefunc"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var AppDB *gorm.DB

func CnnAppDB(dbpath string) {
	var err error
	AppDB, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to %s: %v", dbpath, err)
	}

	if err := AppDB.Exec("PRAGMA foreign_keys = ON;").Error; err != nil {
		log.Fatalf("failed to set foreign_keys pragma: %v", err)
	}

	// Uncomment the following line if you want to enable WAL mode
	// if err := AppDB.Exec("PRAGMA journal_mode=WAL;").Error; err != nil {
	// 	log.Fatalf("failed to set journal_mode pragma: %v", err)
	// }

	SyncAppDB(AppDB)
	log.Println("AppDB connected and synced")
}

func CloseAppDB() {
	sqlDB, err := AppDB.DB()
	if err != nil {
		log.Fatalf("failed to get database instance: %v", err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatalf("failed to close database: %v", err)
	}
	log.Println("AppDB closed")
}

func AppDB_Init() error {
	log.Println("Initializing AppDB defaults")

	sqlFiles := filefunc.GetFileListByExt("static/assets", ".sql")
	for _, sqlFile := range sqlFiles {
		data, err := filefunc.ReadFile(sqlFile)
		if err != nil {
			return err
		}

		if err := AppDB.Exec(string(data)).Error; err != nil {
			return err
		}
	}

	return nil
}
