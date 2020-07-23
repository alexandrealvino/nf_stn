package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type App struct {
	Db *sql.DB
}

func (a *App) Initialize(dbdriver, dbuser, dbpass, dbname string) {
	var err error
	//a.Db, err = sql.Open(dbdriver, dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	a.Db, err = sql.Open(dbdriver, CLEAR_DATABASE_URL)
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
