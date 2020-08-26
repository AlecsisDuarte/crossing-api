package cache

import (
	"crossing-api/libs/log"
)

// Put stores the value in the app's cache with the given key and sets how long it will be stored before expiring
func Put(key string, value interface{}) {
	appCache.SetDefault(key, value)
}

// Get returns the value stored in the app's cache with the given key if found
func Get(key string) (res interface{}, found bool) {
	value, found := appCache.Get(key)
	log.Info("[key %v] not [found %v] in appCache", key, found)
	return value, found
}
