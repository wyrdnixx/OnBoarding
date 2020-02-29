package dbapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"database/sql"	
	"strconv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
)

func ApiItemsHandler(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	if req.Method == "OPTIONS" {
		utils.LogMsg(4, "ApiItemsHandler - HTTP Request : OPTIONS")

	}
	if req.Method == "POST" {
		utils.LogMsg(4, "ApiItemsHandler - HTTP Request : POST")
		DBAddItems(w, req, _appconfig)

	}
	if req.Method == "GET" {
		utils.LogMsg(4, "ApiItemsHandler - HTTP Request : GET")
		DBGetItems(w, _appconfig)
	}
	return nil
}



func DBAddItems(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	utils.LogMsg(4, "############# DBAddItems   ############")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	utils.LogMsg(4, " DBAddItems POST Req.Method Body: "+string(reqBody))



	var data types.NewITM
	errDec := json.Unmarshal(reqBody, &data)

	if errDec != nil {
		utils.LogMsg(1, errDec.Error())
		http.Error(w, errDec.Error(), http.StatusInternalServerError)
		return err 
	} else {
		utils.LogMsg(4, "DBAddItems JSON Unmarshall completed: " + data.Itm.Text)
	}


	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		utils.LogMsg(4, "ERROR: DB Connection: "+ err.Error())
	}
	defer db.Close()

	sql := "INSERT INTO `appdb`.`Items`(`processorId`,`text`,`type`, `enabled`,`depId`) VALUES ('" 	+ 
	strconv.Itoa(data.Itm.ProcessorId) + "', '" + 
	data.Itm.Text + "', '" + 
	data.Itm.ItemType +  "', '1', '" + 
	strconv.Itoa(data.Itm.DepId) +  "');"


	
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


func DBGetItems(w http.ResponseWriter, _appconfig *types.Configuration) error {

	utils.LogMsg(4,"DBGetItems connecting database...")

	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	sql := `SELECT Items.id as id,
	Processors.name as processorId, 
	Items.text as text,
    Items.type as type,
    Items.enabled as enabled,
    Abteilungen.name as depId
	FROM appdb.Items
	left join Abteilungen on Items.depId = Abteilungen.id
    left join Processors on Items.processorId = Processors.id
	
	`

	rows, err := db.Query(sql)
		if err != nil {
			utils.LogMsg(1, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return err
		}

	result := types.Items{}

	for rows.Next() {
		itm := types.Item{}
		err = rows.Scan(&itm.Id,&itm.ProcessorId,&itm.Text, &itm.ItemType, &itm.Enabled, &itm.DepId )
	

	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	result.Items = append(result.Items, itm)
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