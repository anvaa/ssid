package app_menu

import (
	"github.com/spf13/viper"

	"log"
)

var (
	mnuConf         = *viper.New()
	fileName string = "menu.yaml"
	fileType string = "yaml"
	menu     []string
)

func init() {
	mnuConf.SetConfigName(fileName)
	mnuConf.AddConfigPath(".")
	mnuConf.SetConfigType(fileType)
}

func WriteDefaultConfig(appRoot string) {
	// SetDefault sets the default value for the key.
	menu = []string{"Location", "Type", "Manufact", "Serial", "Description", "Price"}
	mnuConf.SetDefault("menutitle", menu)

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

func UpdMenuTitle(key int, value any) {
	menu[key] = value.(string)
	mnuConf.Set("menutitle", menu)
	err := mnuConf.WriteConfig()
	if err != nil {
		log.Fatal("Error writing", fileName)
	}
}

func GetMenuTitles() []string {
	menu = mnuConf.GetStringSlice("menutitle")
	return menu
}
