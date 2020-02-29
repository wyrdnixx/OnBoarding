package utils

import (
	"net/http"
	"fmt"
	"github.com/tkanos/gonfig"
	"github.com/wyrdnixx/go-server/src/go-server/types"
)

func EnableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization,X-CSRF-Token"
	(*w).Header().Set("Access-Control-Allow-Headers", allowedHeaders)
}


////////////////////////////////////////////////////////////////////////
// Level	Meaning
// 0		Critical error messages, aborts the subsystem
// 1		Major error messages, some lost functionality
// 2		Warning error messages which do not cause a functional failure
// 3		Informational messages, showing completion, progress, etc.
// 4		Debug messages, to help in diagnosing a problem
// 5		Trace messages (enter/exit subroutine, buffer contents, etc.)

func LogMsg(_loglvl int, _msg string) {

	// Debug
	if _loglvl == 4 {
		fmt.Println("DEBUG: " + _msg)
	}

	// Info
	if _loglvl == 3 {
		fmt.Println("INFO: " + _msg)
	}
	
	// Warning
	if _loglvl == 2 {
		fmt.Println("WARNING: " + _msg)
	}
	
	// Major
	if _loglvl == 1 {
		fmt.Println("ERROR (Major): " + _msg)
	}
	
	// Critical
	if _loglvl == 0 {
		fmt.Println("ERROR (Critical): " + _msg)
	}
}



var AppConfig = types.Configuration{}
// ReadConfig Lies die Config aus der Config File aus
func ReadConfig() types.Configuration {

	// Config aus file laden:

	AppConfig.Info = "TestInfoComment"

	err := gonfig.GetConf("./config.development.json", &AppConfig)
	if err != nil {
		fmt.Println("ERROR: Config konnte nicht geladen werden: ", err.Error())
	} else {
		fmt.Println("INFO: Config wurde gelesen...")
	}
	return AppConfig
}