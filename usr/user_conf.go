package users

import (

	"log"

	"github.com/spf13/viper"
)


var UsrConf = *viper.New()

var fileName string = "usr.yaml"
var fileType string = "yaml"

func init() {
	UsrConf.SetConfigName(fileName)
	UsrConf.AddConfigPath(".")
	UsrConf.SetConfigType(fileType)
}

func user_WriteConfigFile(app_path string) {

	UsrConf.SetDefault("app_dir", app_path)
	UsrConf.SetDefault("users_db", "data/users.db")
	UsrConf.SetDefault("access_time", "3600*12*1") // 12 hours

	err := UsrConf.WriteConfigAs(fileName)
	if err != nil {
		log.Fatal("Error creating ", fileName)
	}
}

func User_ReadConfig() {
	err := UsrConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading ", fileName)
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
	err := UsrConf.WriteConfigAs(fileName)
	if err != nil {
		log.Fatal("Error SetVal", fileName)
	}
}