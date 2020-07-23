package config

import "os"

//var dbDriver = os.Getenv("MYSQL_DRIVER")
//var dbName = os.Getenv("MYSQL_DATABASE")
//var dbUser = os.Getenv("MYSQL_USER")
//var dbPass = os.Getenv("MYSQL_PASSWORD")
//
var Dbdriver = "mysql"
var Dbuser = "root"
var Dbpass = "admin"
var Dbname = "nf_stn"

var port = os.Getenv("PORT")
var dns = os.Getenv("DATABASE_URL")

//CLEAR_DATABASE_URL = mysql://b4d9a89ec98c72:8222a64a@us-cdbr-east-02.cleardb.com/heroku_256cb7af530bbcb?reconnect=true
var CLEAR_DATABASE_URL = "b4d9a89ec98c72"+":"+"8222a64a"+"@tcp(us-cdbr-east-02.cleardb.com)/"+"heroku_256cb7af530bbcb"