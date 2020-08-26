package cache

import (
	"crossing-api/libs/log"
	"crossing-api/utils"

	"github.com/patrickmn/go-cache"
)

var appCache *cache.Cache

func init() {
	expirationTime := utils.GetCacheExpiration()
	cleanupInterval := utils.GetCacheCleanupInterval()
	log.Info("Initializing the app cache with [expirationTime %v] and [cleanupInterval %v]", expirationTime, cleanupInterval)
	appCache = cache.New(expirationTime, cleanupInterval)
}
