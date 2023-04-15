package main

import (
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,   //use default DB
		PoolSize: 100, //连接池大小
	})
	_, err = rdb.Ping().Result()
	return err
}
