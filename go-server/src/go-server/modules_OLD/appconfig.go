package modules

import (
	"fmt"

	"github.com/tkanos/gonfig"
)

// Configuration -  Allgemeine Config
type Configuration struct {
	ApiPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	Info       string
}

var AppConfig = Configuration{}

// ReadConfig Lies die Config aus der Config File aus
func ReadConfig() {

	// Config aus file laden:

	AppConfig.Info = "TestInfoComment"

	err := gonfig.GetConf("./config.development.json", &AppConfig)
	if err != nil {
		fmt.Println("ERROR: Config konnte nicht geladen werden: ", err.Error())
	} else {
		fmt.Println("INFO: Config wurde gelesen...")
	}

}
