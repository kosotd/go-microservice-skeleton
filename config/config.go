package config

import (
	"go-microservice-skeleton/utils"
	"time"
)

var conf *Config

type Config struct {
	ServerPort        string
	CacheExpiration   string
	CacheUpdatePeriod string
}

func GetConfig() *Config {
	return conf
}

func (c *Config) GetCacheExpiration() time.Duration {
	duration, err := time.ParseDuration(c.CacheExpiration)
	utils.FailOnError(err, "error read cache expiration from config")
	return duration
}

func (c *Config) GetCacheUpdatePeriod() time.Duration {
	duration, err := time.ParseDuration(c.CacheUpdatePeriod)
	utils.FailOnError(err, "error read cache update period from config")
	return duration
}
