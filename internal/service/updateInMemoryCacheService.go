package service

import (
	log "github.com/sirupsen/logrus"
	"github.com/vaibhavahuja/load-route-optimisation/internal/entities"
)

func (app *Application) UpdateValueInMemoryCache(request entities.UpdateCacheRequest) (response entities.UpdateCacheResponse, err error) {
	log.Infof("received request with requestId %s to update cache value to %s", request.RequestId, request.CacheNewValue)
	currentVal := app.myCache
	log.Info("current value of cache is ", currentVal)
	app.myCache = request.CacheNewValue
	response.CacheOldValue = currentVal
	response.CacheNewValue = app.myCache
	response.RequestId = request.RequestId
	return
}
