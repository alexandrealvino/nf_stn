package config

import "os"

// DataBaseConfig interface
type DataBaseConfig interface {
	DbDriver() string
	DbUser() string
	DbPass() string
	DbName() string
	URL() string
	Conn() string
}

// Config struct
type Config struct{}

// DbDriver returns the db driver env variable
func (c *Config) DbDriver() string {
	return os.Getenv("MYSQL_DRIVER")
}

// DbUser returns the db user env variable
func (c *Config) DbUser() string {
	return os.Getenv("MYSQL_USER")
}

// DbPass returns the db password env variable
func (c *Config) DbPass() string {
	return os.Getenv("MYSQL_PASSWORD")
}

// DbName returns the db name env variable
func (c *Config) DbName() string {
	return os.Getenv("MYSQL_DATABASE")
}

// URL returns the URL env variable
func (c *Config) URL() string {
	return os.Getenv("URL")
}

// Conn returns the connection string
func (c *Config) Conn() string {
	conn := c.DbUser() + ":" + c.DbPass() + "@tcp(" + c.URL() + ")/" + c.DbName()
	return conn
}

// RedisConfig interface
type RedisConfig interface {
	DSN() string
}

// RedisCfg struct
type RedisCfg struct{}

// DSN returns the db name env variable
func (r *RedisCfg) DSN() string {
	return os.Getenv("REDIS_DSN")
}
