package helper

import (
	"time"

	"github.com/allegro/bigcache"
)

var Cache, _ = bigcache.NewBigCache(bigcache.DefaultConfig(10 * time.Minute))

func SetCache(key string, value string) bool {
	Cache.Set(key, []byte(value))
	return true
}

func GetCache(key string) string {
	entry, _ := Cache.Get(key)
	return string(entry)
}
