package redis

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initRedis() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		Password:     "",
		DB:           0,
		MinIdleConns: 100,
	})
	_, err = rdb.Ping().Result()
	return err
}
