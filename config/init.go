package config

import (
	"encoding/json"
	"flag"
	"github.com/joho/godotenv"
	"go-microservice-skeleton/utils"
	"os"
	"sync"
)

var once sync.Once

type configGetter interface {
	GetBaseConfig() *Config
}

func InitConfig(configGetter configGetter, loadEnvChild func(EnvHelper)) {
	once.Do(func() {
		conf = configGetter.GetBaseConfig()
		fileName := flag.String("config", "", "Full path to config file")
		flag.Parse()

		if utils.IsNotEmpty(*fileName) {
			loadFileConfiguration(*fileName, conf)
			loadFileConfiguration(*fileName, configGetter)
		} else {
			helper := &envHelperImpl{}
			loadEnvConfiguration(conf, helper)
			loadEnvChild(helper)
		}
	})
}

func loadFileConfiguration(file string, conf interface{}) {
	configFile, err := os.Open(file)
	utils.FailOnError(err, "error open config file")
	defer utils.CloseSafe(configFile)
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&conf)
	utils.FailOnError(err, "error decode config json")
}

func loadEnvConfiguration(conf *Config, helper EnvHelper) {
	_ = godotenv.Load("./config.env")
	conf.ServerPort = helper.GetEnvString(serverPortEnvKey, "8081")
	conf.CacheExpiration = helper.GetEnvString(cacheExpirationEnvKey, "5m")
	conf.CacheUpdatePeriod = helper.GetEnvString(cacheUpdatePeriodEnvKey, "1m")
}
