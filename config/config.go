package config

import "github.com/vaughan0/go-ini"

var file ini.File
var err Error

type Config struct {
	DatabaseName string
	DatabaseUser string
	RedisSecret  string
}

func NewConfig() Config {
	file, err = ini.LoadFile("modip.conf")

	if err != nil {
		panic(err)
	}

	c := Config{}
	c.DatabaseName = get("database_name", "database")
	c.DatabaseUser = get("database_user", "database")
	c.RedisSecret = get("redis_cookie_secret", "redis")

	return c
}

func get(section string, key string) string {
	value, ok := file.Get(key, section)
	if !ok {
		panic("'%s' variable missing from '%s' section", key, section)
	}

	return value
}

