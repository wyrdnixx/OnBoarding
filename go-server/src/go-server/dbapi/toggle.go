package dbapi

import (
	"fmt"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
	 "database/sql"	
	_ "github.com/go-sql-driver/mysql"
	"net/http"		
	"encoding/json"
	"io/ioutil"
	"bytes"
	"strconv"

)



func ApiToggleHandler(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	if req.Method == "OPTIONS" {
		utils.LogMsg(2,"ApiToggleHandler - HTTP Request : OPTIONS - Wrong method")		
	}
	if req.Method == "POST" {
		utils.LogMsg(4,"ApiToggleHandler - HTTP Request : POST")
		toggleEntry(w,req,_appconfig)
	}
	if req.Method == "GET" {
		utils.LogMsg(2,"ApiToggleHandler - HTTP Request : GET - Wrong method")
		
	}
	return nil
}

func toggleEntry (w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {
	reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err
		}
		fmt.Printf("INFO: POST Req.Method Body:  %s \n", reqBody)

		r := bytes.NewReader(reqBody)

		decoder := json.NewDecoder(r)

		var data types.ToggleEntry
		errDec := decoder.Decode(&data)
		if errDec != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err
		}	
			
		utils.LogMsg(4," toggleEntry JSON Got to toggle : " + data.Table + " - ID: " + strconv.Itoa(data.Id))

		var toggleValue int

		if data.Enabled == 1 {
			toggleValue = 0
		}else {
			toggleValue = 1
		}

		db, err := sql.Open("mysql", _appconfig.DBUser + ":" + _appconfig.DBPassword +"@tcp("+_appconfig.DBHost + ":" + _appconfig.DBPort + ")/"+ _appconfig.DBName )
	
		if err != nil {
			fmt.Println("ERROR: DB Connection: ", err.Error())
		}
		defer db.Close()

		var sql string = "update `appdb`."+ data.Table + " set `enabled` = '"+ strconv.Itoa(toggleValue) +"' where id = " +  strconv.Itoa(data.Id) +";"
		utils.LogMsg(4,"Updating: " + sql)
		  insert, err := db.Query(sql)
  		// if there is an error inserting, handle it
  		if err != nil {
			  utils.LogMsg(1,err.Error())			  
			  http.Error (w, err.Error(), http.StatusInternalServerError)
			  return err
  		}
		defer insert.Close()







		w.WriteHeader(http.StatusCreated)
		return nil
}