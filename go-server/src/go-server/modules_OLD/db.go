package modules

import (
	"fmt"
	"database/sql"	
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"encoding/json"
	"log"
	"strconv"
)

func HelloWorld() {
	fmt.Println("Hello")

	//dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		//AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)

	fmt.Printf("INFO: Host=%s Port=%s User=%s DBPasswd=%s DBName=%s\n",
		AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)
	CheckDB()

}

// type Firma struct {
//     Id   int    `json:"id"`
//     Name string `json:"name"`
// }

type Firma struct {
	Id   int `json:"Id"`
	Name string `json:"Name"`
}
type Firmen struct {
	Firmen   []Firma `json:"Firmen"`	
}



func CheckDB () {

	fmt.Println("INFO: Checking DB Connection...")
	// dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName)
	// db, err := sql.Open("mysql", dbInfo)
	//db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	
	db, err := sql.Open("mysql", AppConfig.DBUser + ":" + AppConfig.DBPassword +"@tcp("+AppConfig.DBHost + ":" + AppConfig.DBPort + ")/"+ AppConfig.DBName )

	
		if err != nil {
			fmt.Println("ERROR: DB Connection: ", err.Error())
		}
		defer db.Close()


		results, err := db.Query("Select id, name from Firmen")
		if err != nil {panic(err.Error())}

		for results.Next() {
			var _firma Firma
			err = results.Scan(&_firma.Id, &_firma.Name)
			if err != nil {panic(err.Error())}
			fmt.Println(_firma.Name)
		}
}

func DBgetFirmen  (w http.ResponseWriter) error {

	

	fmt.Println("INFO: Checking DB Connection...")
	db, err := sql.Open("mysql", AppConfig.DBUser + ":" + AppConfig.DBPassword +"@tcp("+AppConfig.DBHost + ":" + AppConfig.DBPort + ")/"+ AppConfig.DBName )

	
		if err != nil {
			fmt.Println("ERROR: DB Connection: ", err.Error())
		}
		defer db.Close()


		rows, err := db.Query("Select id, name from Firmen")
		if err != nil {panic(err.Error())}

		result := Firmen{}
		for rows.Next() {			 
			firma := Firma{}
			err = rows.Scan(&firma.Id, &firma.Name)
	
			if err != nil {
				// handle this error
				panic(err)
			}		
			result.Firmen = append(result.Firmen, firma)
		}

		 js, err := json.Marshal(result)		
		 if err != nil {
			return err 	
		 }
		 fmt.Println("INFO: Sending JSON Result: \n\t " + string(js))	

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	// get any error encountered
	 
	 
	return err

		
}

func DbInsertTest(w http.ResponseWriter) error {
//func DbInsertTest() error {


	///  SQL Transaction
	// https://pseudomuto.com/2018/01/clean-sql-transactions-in-golang/
	fmt.Println("INFO: inserting test SQL Transaction")
	db, err := sql.Open("mysql", AppConfig.DBUser + ":" + AppConfig.DBPassword +"@tcp("+AppConfig.DBHost + ":" + AppConfig.DBPort + ")/"+ AppConfig.DBName )
	handleError(err)
	defer db.Close()

	tx, err := db.Begin()
	handleError(err)

	res, err := tx.Exec("insert into Firmen (name,enabled) values ('Test-Frima',0)")
	if err !=nil {
		tx.Rollback()
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	handleError(err)

	fmt.Println("INFO: Last inserted: ", id)
	handleError(err)

	var sqlQuery string = `insert into Abteilungen (name,firmaId, notifyMail,enabled) values ("TestAbt-2",` +  strconv.FormatInt(id,10) + `,"test@",0)`
	//var sqlQuery = strconv.FormatInt(id,10)
	handleError(err)
	fmt.Println("INFO: SQL Statement: "+ sqlQuery)
	res,err = tx.Exec(sqlQuery)
	if err !=nil {		
		fmt.Println("ERROR: ", err)
		//w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))	
		tx.Rollback()
		//log.Fatal(err)
		
	}else {
		handleError(tx.Commit())
	}
	
	log.Println("INFO: SQL Transaction done...")
	
	//w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("☄ HTTP status code returned!"))

	return err
}


func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}