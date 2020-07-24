package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql" // importing driver mysql
	"log"
)

// App class object for db instantiation
type App struct {
	Db *sql.DB
}

// Initialize initiates the db connection
func (a *App) Initialize(Dbdriver, Dbuser, Dbpass, Dbname string) {
	var err error
	//a.Db, err = sql.Open(Dbdriver, Dbuser+":"+Dbpass+"@tcp(127.0.0.1:3306)/"+Dbname)
	a.Db, err = sql.Open("mysql", "root"+":"+"admin"+"@tcp(mysql:3306)/"+"nf_stn")
	//a.Db, err = sql.Open(dbdriver, CLEAR_DATABASE_URL)
	if err != nil {
		log.Fatal(err)
	}
}

//func  DbConn() (db *sql.DB) {
//	//dbdriver := dbDriver
//	//dbuser := dbUser
//	//dbpass := dbPass
//	//dbname := dbName
//	dbdriver := "mysql"
//	dbuser := "root"
//	dbpass := "admin"
//	dbname := "nf_stn"
//	////dbpass := "!Q2w#E4r"
//	db, err := sql.Open(dbdriver, dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
//	if err != nil {
//		panic(err.Error())
//	}
//	return db
//}
