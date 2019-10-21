package server

import (
	"encoding/json"
	"github.com/pkg/errors"
	"go-microservice-skeleton/cache"
	"go-microservice-skeleton/config"
	"go-microservice-skeleton/server"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http/httptest"
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

func TestCacheAndWrite(t *testing.T) {
	recorder := httptest.NewRecorder()

	err := server.CacheAndWrite(recorder, "cache", func() (resp interface{}, err error) {
		return nil, errors.New("error1")
	})
	assert.Error(t, err, "server.CacheAndWrite -> responseSupplier(): error1")

	err = server.CacheAndWrite(recorder, "cache", func() (resp interface{}, err error) {
		return map[string]string{"key": "value"}, nil
	})
	assert.NilError(t, err)

	body, err := ioutil.ReadAll(recorder.Body)
	assert.NilError(t, err)

	res := map[string]string{}
	err = json.Unmarshal(body, &res)
	assert.NilError(t, err)

	assert.Equal(t, res["key"], "value")

	err = server.CacheAndWrite(recorder, "cache", func() (resp interface{}, err error) {
		return nil, errors.New("error1")
	})
	assert.NilError(t, err)

	time.Sleep(1 * time.Second)

	err = server.CacheAndWrite(recorder, "cache", func() (resp interface{}, err error) {
		return nil, errors.New("error1")
	})
	assert.Error(t, err, "server.CacheAndWrite -> responseSupplier(): error1")
}
