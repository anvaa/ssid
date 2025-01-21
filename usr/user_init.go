package users

import (
	"srv/filefunc"
	"srv/srv_conf"
	"usr/user_embed"

	"log"
)

// User_Init initializes the user database and configuration file
func UserInit(app_path string) error {

	// user embed files
	err := user_embed.User_EmbedFiles()
	if err != nil {
		return err
	}

	user_Database(srv_conf.DataDir)

	// Check/make usr.yaml
	if !filefunc.IsExists(app_path + "/usr.yaml") {
		WriteConfigFile(app_path)
	}
	ReadConfig()

	return nil
}

func user_Database(dataFolder string) {
	// Check for data folder and users.db file
	userdb := dataFolder + "/users.db"
	if !filefunc.IsExists(dataFolder) {
		log.Println("No data folder found. Creating", dataFolder)
		filefunc.CreateFolder(dataFolder)
		filefunc.CreateFile(userdb)
	}

	// connect/sync to the database
	CnnUserDB(userdb)
}
