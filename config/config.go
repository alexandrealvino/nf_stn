package config

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql" // importing driver mysql
	"log"
	"os"
)

var  Client *redis.Client

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

	//Initializing redis
	dsn := os.Getenv("REDIS_DSN") // TODO env
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	Client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err = Client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
