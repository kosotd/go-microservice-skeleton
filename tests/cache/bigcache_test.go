package cache

import (
	"go-microservice-skeleton/cache"
	"go-microservice-skeleton/config"
	"gotest.tools/assert"
	"testing"
	"time"
)

type testConfig struct {
	config config.Config
}

func (c *testConfig) GetBaseConfig() *config.Config {
	return &c.config
}

func init() {
	conf := &testConfig{}
	config.InitConfig(conf, func(helper config.EnvHelper) {})
	conf.GetBaseConfig().CacheExpiration = "200ms"
	conf.GetBaseConfig().CacheUpdatePeriod = "1ms"
	cache.InitBigCache()
}

func TestBigCache(t *testing.T) {
	err := cache.SetData("key", []byte("value"))
	assert.NilError(t, err)

	value, ok := cache.GetData("key")
	assert.Equal(t, ok, true)
	assert.Equal(t, "value", string(value))

	time.Sleep(1 * time.Second)
	_, ok = cache.GetData("key")
	assert.Equal(t, ok, false)
}
