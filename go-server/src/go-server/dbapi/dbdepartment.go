package dbapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/wyrdnixx/go-server/src/go-server/types"
	"github.com/wyrdnixx/go-server/src/go-server/utils"
	// "bytes"
)

func ApiDepartmentsHandler(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	if req.Method == "OPTIONS" {
		utils.LogMsg(4, "ApiDepartmentsHandler - HTTP Request : OPTIONS")

	}
	if req.Method == "POST" {
		utils.LogMsg(4, "ApiDepartmentsHandler - HTTP Request : POST")
		DBAddDepartment(w, req, _appconfig)

	}
	if req.Method == "GET" {
		utils.LogMsg(4, "ApiDepartmentsHandler - HTTP Request : GET")
		DBGetDepartments(w, _appconfig)
	}
	return nil
}

func DBAddDepartment(w http.ResponseWriter, req *http.Request, _appconfig *types.Configuration) error {

	//	utils.EnableCors(&w)
	utils.LogMsg(4, "############ DBAddDepartment  ######### ")

	reqBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	utils.LogMsg(4, "############ DBAddDepartment  ######### ")
	utils.LogMsg(4, " DBAddDepartment POST Req.Method Body: "+string(reqBody))

	var data types.NewDep
	errDec := json.Unmarshal(reqBody, &data)

	///////////////////////////////////////   AAAhhhhhhh..... ToDo...

	if errDec != nil {
		utils.LogMsg(1, errDec.Error())
		http.Error(w, errDec.Error(), http.StatusInternalServerError)
		return err
	}

	utils.LogMsg(3, "AddDepartment JSON got newDepNAme: "+data.Dep.NewDepName)

	/// DB insert
	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		fmt.Println("ERROR: DB Connection: ", err.Error())
	}
	defer db.Close()

	sql := "INSERT INTO `appdb`.`Abteilungen`(`name`,`FirmaId`,`notifyMail`, `enabled`) VALUES ('" + data.Dep.NewDepName + "', '" + strconv.Itoa(data.Dep.NewDepCompany) + "', '" + data.Dep.NewDepMail + "','1');"
	//	sql := `TEST`

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

func DBGetDepartments(w http.ResponseWriter, _appconfig *types.Configuration) error {

	fmt.Println("INFO: DBGetDepartments: Checking DB Connection...")
	fmt.Println("INFO: DBGetDepartments: Appconfig: " + _appconfig.DBUser + " - " + _appconfig.DBHost)

	db, err := sql.Open("mysql", _appconfig.DBUser+":"+_appconfig.DBPassword+"@tcp("+_appconfig.DBHost+":"+_appconfig.DBPort+")/"+_appconfig.DBName)

	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	defer db.Close()

	sql := `select Abteilungen.id as id, Abteilungen.name as name , Firmen.name as firma, notifyMail, Abteilungen.enabled 
		from Abteilungen 
		left join Firmen on 
		Abteilungen.firmaId = Firmen.id`

	utils.LogMsg(4, "DB Query: "+sql)
 
	rows, err := db.Query(sql)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}

	result := types.Departments{}
	for rows.Next() {
		dep := types.Department{}
		err = rows.Scan(&dep.Id, &dep.Name, &dep.Firma, &dep.NotifyMail, &dep.Enabled)

		if err != nil {
			utils.LogMsg(1, err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		result.Departments = append(result.Departments, dep)
	}

	js, err := json.Marshal(result)
	if err != nil {
		utils.LogMsg(1, err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return err
	}
	fmt.Println("INFO: Sending JSON Result: \n\t " + string(js))

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(js)
	return err

}
