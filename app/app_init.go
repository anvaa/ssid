package app

import (
	"log"

	"app/app_conf"
	"app/app_db"
	"app/app_embed"

	"srv/filefunc"
	"srv/srv_conf"
)

func AppInit(app_folder string) error {

	// app embed files
	err := app_embed.App_EmbedFiles()
	if err != nil {
		return err
	}

	// check for app config file
	configFile := app_folder + "/app.yaml"
	if !filefunc.IsExists(configFile) {
		log.Println("No app.yaml file found. Creating", configFile)
		app_conf.WriteDefaultConfig(app_folder)
	}
	app_conf.ReadConfig() // read the config file


	var isNew bool = false
	appdb := srv_conf.DataDir + "/app.db"
	if !filefunc.IsExists(appdb) {
		log.Println("Creating", appdb)
		filefunc.CreateFile(appdb)
		isNew = true
		
	}

	// connect/sync to the app database
	app_db.CnnAppDB(appdb)

	if isNew {
		err := app_db.AppDB_Init()
		if err != nil {
			return err
		}
	}

	return nil
}

