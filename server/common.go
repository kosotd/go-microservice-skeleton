package server

import (
	"bytes"
	"encoding/json"
	"github.com/pkg/errors"
	"go-microservice-skeleton/cache"
	"net/http"
)

func CacheAndWrite(w http.ResponseWriter, cacheName string, responseSupplier func() (interface{}, error)) error {
	var data []byte
	var has bool
	if data, has = cache.GetData(cacheName); !has {
		resp, err := responseSupplier()
		if err != nil {
			return errors.Wrapf(err, "server.CacheAndWrite -> responseSupplier()")
		}
		var buff bytes.Buffer
		err = json.NewEncoder(&buff).Encode(resp)
		if err != nil {
			return errors.Wrapf(err, "server.CacheAndWrite -> json.NewEncoder(&buff).Encode")
		}
		data = buff.Bytes()
		err = cache.SetData(cacheName, data)
		if err != nil {
			return errors.Wrapf(err, "server.CacheAndWrite -> cache.SetData")
		}
	}

	_, err := w.Write(data)
	if err != nil {
		return errors.Wrapf(err, "server.CacheAndWrite -> w.Write")
	}
	return nil
}
