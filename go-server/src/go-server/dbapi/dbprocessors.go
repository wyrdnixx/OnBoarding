package dbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
)

func ApiProcessorsHandler(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	if req.Method == "OPTIONS" {
		utils.LogMsg(4, "ApiProcessorsHandler - HTTP Request : OPTIONS")

	}
	if req.Method == "POST" {
		utils.LogMsg(4, "ApiProcessorsHandler - HTTP Request : POST")
		DBAddDepartment(w, req, _appconfig)

	}
	if req.Method == "GET" {
		utils.LogMsg(4, "ApiProcessorsHandler - HTTP Request : GET")
		DBGetDepartments(w, _appconfig)
	}
	return nil
}

func DBAddProcessors(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	utils.LogMsg(4, "############# DBAddProcessor   ############")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	var data types.NewProc
	errDec := json.Unmarshal(reqBody, &data)

	if errDec != nil {
		utils.LogMsg(1, errDec.Error())
		http.Error(w, errDec.Error(), http.StatusInternalServerError)
		return err
	}

	utils.LogMsg(3, "DBAddProcessor JSON got newDepNAme: "+data.NewProcName)

	/// Write http Status
	w.WriteHeader(http.StatusCreated)
	return nil

}
