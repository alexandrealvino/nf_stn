package config

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
	_ "github.com/go-sql-driver/mysql" // importing driver mysql
	log "github.com/sirupsen/logrus"
)

// App class object for db instantiation
type App struct {
	Db  *sql.DB
	Clt *redis.Client
}

// Initialize initiates the db connection and redis
func (a *App) Initialize(dbDriver, conn string) {
	var err error
	// Initializing db connection
	a.Db, err = sql.Open(dbDriver, conn)
	if err != nil {
		log.Fatal(err)
	}
}

// ConnectRedis start a connection with redis
func (a *App) ConnectRedis(dsn string) {
	// Initializing redis
	a.Clt = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := a.Clt.Ping().Result()
	if err != nil {
		panic(err)
	}
	log.Info("redis client connected")
}

//
