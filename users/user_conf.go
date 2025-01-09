package users

import (

	"log"

	"github.com/spf13/viper"
)


var UsrConf = viper.New()

var usrFileName string = "usr.yaml"

func user_WriteConfigFile(app_path string) {

	UsrConf.SetDefault("app_dir", app_path)
	UsrConf.SetDefault("users_db", "data/users.db")
	
	UsrConf.SetDefault("access_time", "3600*12*1") // 12 hours

	err := UsrConf.WriteConfigAs(usrFileName)
	if err != nil {
		log.Fatal("Error creating", usrFileName)
	}
}

func User_ReadConfig() {
	err := UsrConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading", usrFileName)
	}
}

func GetString(key string) string {
	return UsrConf.GetString(key)
}

func GetInt(key string) int {
	return UsrConf.GetInt(key)
}

func GetInt64(key string) int64 {
	return UsrConf.GetInt64(key)
}

func GetBool(key string) bool {
	return UsrConf.GetBool(key)
}

func SetVal(key string, val any) {
	UsrConf.Set(key, val)
	err := UsrConf.WriteConfigAs(usrFileName)
	if err != nil {
		log.Fatal("Error SetVal", usrFileName)
	}
}