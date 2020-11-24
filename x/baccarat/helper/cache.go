package helper

import "github.com/allegro/bigcache"
import "time"
import "fmt"

var Cache *bigcache.BigCache

func init() {
	var err error
	Cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(30 * time.Minute))
	if err != nil {
		fmt.Println("Couldn't init cache: \n%w\n", err.Error())
	}
}

func SetCache(key string, value string) bool {
	err := Cache.Set(key, []byte(value))
	if err != nil {
		fmt.Println("Couldn't set cache: \n%w\n", err.Error())
		return false
	}
	return true
}

func GetCache(key string) (string, bool) {
	entry, err := Cache.Get(key)
	if err != nil {
		fmt.Println("Couldn't get cache: \n%w\n", err.Error())
		return string(entry), false
	}
	return string(entry), true
}
