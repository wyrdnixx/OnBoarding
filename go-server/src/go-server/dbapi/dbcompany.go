package dbapi

import (
	"fmt"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
	"database/sql"	
	_"github.com/go-sql-driver/mysql"
	"net/http"		
	"encoding/json"
	"io/ioutil"
	"bytes"
	// "strconv"

)

func ApiCompanysHandler(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	if req.Method == "OPTIONS" {
		utils.LogMsg(4,"ApiCompanysHandler - HTTP Request : OPTIONS")
		
	}
	if req.Method == "POST" {
		utils.LogMsg(4,"ApiCompanysHandler - HTTP Request : POST")
		DBAddCompany(w, req,_appconfig)
		
	}
	if req.Method == "GET" {
		utils.LogMsg(4,"ApiCompanysHandler - HTTP Request : GET")
		DBGetCompanys(w,_appconfig)
	}
	return nil
}

func DBAddCompany (w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

//	utils.EnableCors(&w)
	
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err
		}
		fmt.Printf("INFO: POST Req.Method Body:  %s \n", reqBody)

		r := bytes.NewReader(reqBody)

		decoder := json.NewDecoder(r)

		var data types.NewCompany
		errDec := decoder.Decode(&data)
		if errDec != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err
		}	
			
		utils.LogMsg(3,"AddCompany JSON Got: " + data.Name )
	

		/// DB insert
		db, err := sql.Open("mysql", _appconfig.DBUser + ":" + _appconfig.DBPassword +"@tcp("+_appconfig.DBHost + ":" + _appconfig.DBPort + ")/"+ _appconfig.DBName )
	
		if err != nil {
			fmt.Println("ERROR: DB Connection: ", err.Error())
		}
		defer db.Close()
		
  		// perform a db.Query insert
		var sql string = "INSERT INTO `appdb`.`Firmen` (`name`, `enabled`) VALUES ('" + data.Name +"', '1');"
		utils.LogMsg(4,"Updating: " + sql)
		  insert, err := db.Query(sql)
  		// if there is an error inserting, handle it
  		if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			  return err
  		}
		defer insert.Close()


		/// Write http Status
	w.WriteHeader(http.StatusCreated)
	return nil
}

func DBGetCompanys  (w http.ResponseWriter, _appconfig *types.Configuration) error {

	fmt.Println("INFO: DBGetCompanys: Checking DB Connection...")
	fmt.Println("INFO: DBGetCompanys: Appconfig: " + _appconfig.DBUser + " - " + _appconfig.DBHost)

	db, err := sql.Open("mysql", _appconfig.DBUser + ":" + _appconfig.DBPassword +"@tcp("+_appconfig.DBHost + ":" + _appconfig.DBPort + ")/"+ _appconfig.DBName )
	
		if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err 
		}
		defer db.Close()
		var sql string = "Select id, name, enabled from Firmen"
		utils.LogMsg(4,"DB Query: " + sql)
		rows, err := db.Query(sql)
		if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err 
		}

		result := types.Firmen{}
		for rows.Next() {			 
			firma :=  types.Firma{}
			err = rows.Scan(&firma.Id, &firma.Name, &firma.Enabled)
	
			if err != nil {
				utils.LogMsg(1,err.Error())			  
				http.Error (w, err.Error(), http.StatusInternalServerError)
			}		
			result.Firmen = append(result.Firmen, firma)
		}

		 js, err := json.Marshal(result)		
		 if err != nil {
			utils.LogMsg(1,err.Error())			  
			http.Error (w, err.Error(), http.StatusInternalServerError)
			return err 	
		 }
		 fmt.Println("INFO: Sending JSON Result: \n\t " + string(js))	

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	return err
		
}

