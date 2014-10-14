package config

import "github.com/vaughan0/go-ini"

var file ini.File
var err error

type Config struct {
	DatabaseName     string
	DatabaseUser     string
	DatabaseHost     string
	DatabaseUserPass string
	RedisSecret      []byte
	RedisPort         string
}

func (c *Config) PgConnectionString() string {
	conn := "postgres://"
	conn += c.DatabaseUser
	conn += ":"

	if c.DatabaseUserPass != "" {
		conn += c.DatabaseUserPass
	}

	conn += "@"
	conn += c.DatabaseHost + "/"
	conn += c.DatabaseName
	conn += "?sslmode=disable"
	return conn
}

func NewConfig() Config {
	file, err = ini.LoadFile("modip.conf")

	if err != nil {
		panic(err)
	}

	c := Config{}
	c.DatabaseName = get("database_name", "database")
	c.DatabaseUser = get("database_user", "database")
	c.DatabaseHost = get("host", "database")
	c.DatabaseUserPass = get("user_password", "database")
	c.RedisSecret = []byte(get("redis_cookie_secret", "redis"))
	c.RedisPort = ":" + get("port", "redis")

	return c
}

func get(section string, key string) string {
	value, ok := file.Get(key, section)
	if !ok {
		panic("The Config file is wrong")
	}

	return value
}

