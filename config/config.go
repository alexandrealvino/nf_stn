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
func (a *App) Initialize(dbdriver, dbuser, dbpass, dbname string) {
	var err error
	a.Db, err = sql.Open(dbdriver, dbuser+":"+dbpass+"@tcp(127.0.0.1:3306)/"+dbname)
	if err != nil {
		log.Fatal(err)
	}
}
