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
	mnu := []string{"Location", "Type", "Manufact"}
	mnuConf.SetDefault("menutitle", mnu)

	err := mnuConf.WriteConfigAs(fileName)
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func ReadMenuConfig() {
	err := mnuConf.ReadInConfig()
	if err != nil {
		log.Fatal("Error reading", fileName)
	}
}

func UpdMenuTitle(idx int, title string) {
	mnu[idx] = title
	mnuConf.Set("menutitle", mnu)
	err := mnuConf.WriteConfig()
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func GetMenuTitles() []string {
	mnu = mnuConf.GetStringSlice("menutitle")
	return mnu
}