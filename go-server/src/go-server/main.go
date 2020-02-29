package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wyrdnixx/go-server/src/go-server/dbapi"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
)

var AppConfig = types.Configuration{}

//// Info : HTTP Statuscodes:
//   https://golang.org/pkg/net/http/#pkg-constants
//   w.WriteHeader(http.StatusInternalServerError)
/////

func defaultHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Nothing here... use API links")
	utils.LogMsg(2, "/ DefaultHandler was requested from "+req.RemoteAddr+" \n: "+req.URL.Path+" : Nothing here... Wrong api call...")

}

func toggleHandler(w http.ResponseWriter, req *http.Request) {
	utils.EnableCors(&w)

	//fmt.Fprint(w, "Nothing here... use API links")
	utils.LogMsg(4, req.URL.Path+" was requested from "+req.RemoteAddr)
	dbapi.ApiToggleHandler(w, req, &AppConfig)

}

func companysHandler(w http.ResponseWriter, req *http.Request) {

	utils.EnableCors(&w)

	utils.LogMsg(3, "/api/Companys was requested...")
	//dbapi.DBGetCompanys(w, &AppConfig)
	dbapi.ApiCompanysHandler(w, req, &AppConfig)
}

func departmentsHandler(w http.ResponseWriter, req *http.Request) {

	utils.EnableCors(&w)
	utils.LogMsg(3, "/api/Departments was requested...")
	//dbapi.DBGetCompanys(w, &AppConfig)
	dbapi.ApiDepartmentsHandler(w, req, &AppConfig)
}

func processorsHandler(w http.ResponseWriter, req *http.Request) {

	utils.EnableCors(&w)
	utils.LogMsg(3, "/api/Processors was requested...")
	dbapi.ApiProcessorsHandler(w, req, &AppConfig)
}

func itemsHandler(w http.ResponseWriter, req *http.Request) {

	utils.EnableCors(&w)
	utils.LogMsg(3, "/api/Items was requested...")
	dbapi.ApiItemsHandler(w, req, &AppConfig)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.HandleFunc("/api/Companys", companysHandler)
	http.HandleFunc("/api/Departments", departmentsHandler)
	http.HandleFunc("/api/Processors", processorsHandler)
	http.HandleFunc("/api/Items", itemsHandler)

	http.HandleFunc("/api/toggle", toggleHandler)

	AppConfig = utils.ReadConfig()

	utils.LogMsg(3, " Api Server listening on Port:"+AppConfig.ApiPort)
	log.Fatal(http.ListenAndServe(AppConfig.ApiPort, nil))

}
