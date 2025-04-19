package app_menu

import (
	"github.com/spf13/viper"

	"log"

)

var (
	mnuConf = *viper.New()
	fileName string = "mnu.yaml"
	fileType string = "yaml"
	mnu      []string
)

func init() {
	mnuConf.SetConfigName(fileName)
	mnuConf.AddConfigPath(".")
	mnuConf.SetConfigType(fileType)
}

func WriteDefaultConfig(appRoot string) {
	// SetDefault sets the default value for the key.
	mnuConf.SetDefault("menutitle", []string{"Location", "Type", "Manufact"})

	err := mnuConf.WriteConfigAs(fileName)
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func ReadConfig() {
	err := mnuConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading", fileName)
	}
}

func UpdMenu(key int, value any) {
	mnu[key] = value.(string)
	mnuConf.Set("menutitle", mnu)
	err := mnuConf.WriteConfig()
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func GetMenu() []string {
	mnu = mnuConf.GetStringSlice("menutitle")
	return mnu
}