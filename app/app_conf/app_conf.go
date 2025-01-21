package app_conf

import (
	
	"log"
	"time"

	"github.com/spf13/viper"
)

var appConf = *viper.New()

var fileName string = "app.yaml"
var fileType string = "yaml"

func init() {
	appConf.SetConfigName(fileName)
	appConf.AddConfigPath(".")
	appConf.SetConfigType(fileType)
}

func WriteDefaultConfig(appRoot string) {
	// SetDefault sets the default value for the key.
	appConf.SetDefault("app_db", "data/app.db")
	appConf.SetDefault("txt_len", 35)
	appConf.SetDefault("usr_url", "/app")

	err := appConf.WriteConfigAs(fileName)
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func ReadConfig() {
	err := appConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading", fileName)
	}
}

func GetString(key string) string {
	return appConf.GetString(key)
}

func GetInt(key string) int {
	return appConf.GetInt(key)
}

func GetInt64(key string) int64 {
	return appConf.GetInt64(key)
}

func GetTime(key string) time.Time {
	return appConf.GetTime(key)
}

func GetBool(key string) bool {
	return appConf.GetBool(key)
}

func SetVal(key string, val any) {
	appConf.Set(key, val)
	err := appConf.WriteConfigAs("app.yaml")
	if err != nil {
		log.Fatal("Error SetVal", fileName)
	}
}

func TxtLength() int {
	return appConf.GetInt("txt_len")
}




