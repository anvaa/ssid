package app_db

import (

	"srv/filefunc"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"log"

)

var AppDB *gorm.DB

func CnnAppDB(dbpath string) {
	var err error
	AppDB, err = gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect", dbpath)
	}

	AppDB.Exec("PRAGMA foreign_keys = ON;")
	// AppDB.Exec("PRAGMA journal_mode=WAL;")

	SyncAppDB(AppDB)
	log.Println("AppDB connected and synced")
}

func CloseAppDB() {
	sqlDB, err := AppDB.DB()
	if err != nil {
		log.Fatal("failed to close database")
	}
	sqlDB.Close()
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

		// execute sql file
		err = AppDB.Exec(string(data)).Error
		if err != nil {
			return err
		}
	}
	
	return nil
	
}