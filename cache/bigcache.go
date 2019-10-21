package cache

import (
	"github.com/allegro/bigcache"
	"github.com/kosotd/go-microservice-skeleton/config"
	"github.com/kosotd/go-microservice-skeleton/utils"
	"github.com/pkg/errors"
	"sync"
)

var cache *bigcache.BigCache = nil
var once sync.Once

func InitBigCache() {
	once.Do(func() {
		defaultConfig := bigcache.DefaultConfig(config.GetConfig().GetCacheExpiration())
		defaultConfig.CleanWindow = config.GetConfig().GetCacheUpdatePeriod()
		cache, _ = bigcache.NewBigCache(defaultConfig)
	})
}

func SetData(key string, data []byte) error {
	utils.FailIfNil(cache, "cache are not initialized")

	if err := cache.Set(key, data); err != nil {
		return errors.Wrapf(err, "cache.SetData -> cache.Set(%s)", key)
	}
	return nil
}

func GetData(key string) ([]byte, bool) {
	utils.FailIfNil(cache, "cache are not initialized")

	if data, err := cache.Get(key); err == nil {
		return data, true
	}
	return nil, false
}
