package cache

import (
	"crossing-api/libs/log"

	"github.com/patrickmn/go-cache"
	"time"
)

// Put stores the value in the app's cache with the given key and sets how long it will be stored before expiring
func Put(key string, value interface{}) {
	appCache.SetDefault(key, value)
}

// PutWithDuration stores the value in the app's cache with the given key and marks it as valid as long as the duration provided
func PutWithDuration(key string, value interface{}, duration time.Duration) {
	appCache.Set(key, value, duration)
}

// PutNoExpiration stores the value in the app's cache without an expiration time
func PutNoExpiration(key string, value interface{}) {
	appCache.Set(key, value, cache.NoExpiration)
}

// Get returns the value stored in the app's cache with the given key if found
func Get(key string) (res interface{}, found bool) {
	value, found := appCache.Get(key)
	log.Info("[key %v] [found %v] in appCache", key, found)
	return value, found
}
