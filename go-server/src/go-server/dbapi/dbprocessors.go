package dbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"	
	"fmt"
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
		DBAddProcessors(w, req, _appconfig)

	}
	if req.Method == "GET" {
		utils.LogMsg(4, "ApiProcessorsHandler - HTTP Request : GET")
		DBGetProcessors(w, _appconfig)
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

	utils.LogMsg(4, " DBAddProcessors POST Req.Method Body: "+string(reqBody))
	

	var datax types.NewProc
	errDec := json.Unmarshal(reqBody, &datax)

	if errDec != nil {
		utils.LogMsg(1, errDec.Error())
		http.Error(w, errDec.Error(), http.StatusInternalServerError)
		return err 
	} else {
		utils.LogMsg(4, "Unmarshall completed: " + string(datax.Item.NewProcessorName))
	}

	utils.LogMsg(3, "DBAddProcessor JSON got newProcName: "+datax.Item.NewProcessorName)


	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		fmt.Println("ERROR: DB Connection: ", err.Error())
	}
	defer db.Close()

	sql := "INSERT INTO `appdb`.`Processors`(`name`,`notifyMail`, `enabled`) VALUES ('" + datax.Item.NewProcessorName + "', '" + datax.Item.NewProcessorMail + "','1');"


	utils.LogMsg(4, "Updating: "+sql)
	insert, err := db.Query(sql)
	// if there is an error inserting, handle it
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer insert.Close()



	/// Write http Status
	w.WriteHeader(http.StatusCreated)
	return nil

}


func DBGetProcessors(w http.ResponseWriter, _appconfig *types.Configuration) error {

	utils.LogMsg(4,"DBGetProcessors connecting database...")

	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	sql := `select Processors.id as id, Processors.name as name, Processors.notifyMail as notifyMail, enabled
		from Processors`

	rows, err := db.Query(sql)
		if err != nil {
			utils.LogMsg(1, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

	result := types.Processors{}

	for rows.Next() {
		proc := types.Processor{}
		err = rows.Scan(&proc.Id,&proc.Name,&proc.NotifyMail,&proc.Enabled)
	
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result.Processors = append(result.Processors, proc)
	}


	js, err := json.Marshal(result)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	utils.LogMsg(4,"Sending JSON Result: \n\t " + string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	return err

}