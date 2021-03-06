package config

import (
	"github.com/kosotd/go-microservice-skeleton/utils"
	"github.com/pkg/errors"
	"time"
)

var conf *Config
var initialized int

type Config struct {
	ServerPort        string
	CacheExpiration   string
	CacheUpdatePeriod string
	LogLevel          int
}

func GetConfig() *Config {
	utils.FailIfNotInitialized(initialized, "config not initialized")
	return conf
}

func (c *Config) GetCacheExpiration() time.Duration {
	duration, err := time.ParseDuration(c.CacheExpiration)
	utils.FailIfError(errors.Wrapf(err, "error read cache expiration from config"))
	return duration
}

func (c *Config) GetCacheUpdatePeriod() time.Duration {
	duration, err := time.ParseDuration(c.CacheUpdatePeriod)
	utils.FailIfError(errors.Wrap(err, "error read cache update period from config"))
	return duration
}
