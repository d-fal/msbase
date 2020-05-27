package cache

import (
	"encoding/json"
	"fmt"

	"msbase/pkg/conf"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
	conn redis.Conn
)

func init() {

}

// PrepareRedisPool gets the pool provided by redis
func PrepareRedisPool() {

	pool = &redis.Pool{
		MaxIdle:   conf.GetConfigObject().GetCacheConfig().MaxIdle,
		MaxActive: conf.GetConfigObject().GetCacheConfig().MaxActive,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(conf.GetConfigObject().GetCacheConfig().Proto,
				fmt.Sprintf("%s", conf.GetConfigObject().GetCacheConfig().Address))
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// GetRedisPoolObject returns redis pool
func GetRedisPoolObject() *redis.Pool {
	if &pool == nil {
		PrepareRedisPool()
	}
	return pool
}

// GetRedisConnection returns the redis connection
func GetRedisConnection() *redis.Conn {
	if conn == nil {
		conn = GetRedisPoolObject().Get()
	}
	return &conn
}

// Ping pings cache server
func Ping() error {
	pong, err := conn.Do("PING")
	if err != nil {
		return err
	}

	s, err := redis.String(pong, err)
	if err != nil {
		return err
	}
	fmt.Printf("PING RESPONSE = %s\n", s)
	return nil
}

// SetInCache sets data in cache server
func SetInCache(key string, value interface{}) {
	jsonStr, _ := json.Marshal(value)
	(*GetRedisConnection()).Do("SET", key, jsonStr)
}

// RetrieveCache retrieves cache in the
func RetrieveCache(key string) (interface{}, conf.ErrorBlock) {

	s, err := redis.String((*GetRedisConnection()).Do("GET", key))

	if err == redis.ErrNil {
		return nil, conf.GetConfigObject().GetErrorList().WarningKeyDoesNotExistInCache
	} else if err != nil {
		return nil, conf.GetConfigObject().GetErrorList().ErrorCannotQueryCacheServer
	}

	return []byte(s), conf.ErrorBlock{}
}
