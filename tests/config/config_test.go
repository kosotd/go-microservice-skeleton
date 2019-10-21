package config

import (
	"go-microservice-skeleton/config"
	"gotest.tools/assert"
	"testing"
)

type testConfig struct {
	config config.Config
}

func (c *testConfig) GetBaseConfig() *config.Config {
	return &c.config
}

func TestConfigEnv(t *testing.T) {
	config.InitConfig(&testConfig{}, func(helper config.EnvHelper) {})

	assert.Equal(t, "9090", config.GetConfig().ServerPort)
	assert.Equal(t, "1s", config.GetConfig().CacheExpiration)
	assert.Equal(t, "2s", config.GetConfig().CacheUpdatePeriod)
}
