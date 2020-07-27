package config

import "os"

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