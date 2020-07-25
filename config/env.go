package config

import "os"

//var DbDriver = os.Getenv("MYSQL_DRIVER")
//var DbName = os.Getenv("MYSQL_DATABASE")
//var DbUser = os.Getenv("MYSQL_USER")
//var DbPass = os.Getenv("MYSQL_PASSWORD")
//var DbRootPwd = os.Getenv("MYSQL_ROOT_PASSWORD")
////

type DataBaseConfig interface {
	Dbdriver() string
	Dbuser() string
	Dbpass() string
	Dbname() string
}

type Config struct {}

func (c *Config) Dbdriver() string {
	return os.Getenv("MYSQL_DRIVER")
}

func (c *Config) Dbuser() string {
	return os.Getenv("MYSQL_USER")
}

func (c *Config) Dbpass() string {
	return os.Getenv("MYSQL_PASSWORD")
}
func (c *Config) Dbname() string {
	return os.Getenv("MYSQL_DATABASE")
}

////Dbdriver is the driver name
//var Dbdriver = "mysql"
//
////Dbuser is the username for the db connection
//var Dbuser = "root"
//
////Dbpass is the password for the db connection
//var Dbpass = "admin"
//
////Dbname is the db name
//var Dbname = "nf_stn"

//var port = os.Getenv("PORT")
//var dns = os.Getenv("DATABASE_URL")
//
////CLEAR_DATABASE_URL = mysql://b4d9a89ec98c72:8222a64a@us-cdbr-east-02.cleardb.com/heroku_256cb7af530bbcb?reconnect=true
//var CLEAR_DATABASE_URL = "b4d9a89ec98c72"+":"+"8222a64a"+"@tcp(us-cdbr-east-02.cleardb.com)/"+"heroku_256cb7af530bbcb"
