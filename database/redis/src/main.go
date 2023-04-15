package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.Client

func InitRDB() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	_ = rdb.Ping()
}

func main() {
	//ctx := context.Background()
	InitRDB()
	err := rdb.Set("testKey", "testValue", 0).Err()
	if err != nil {
		fmt.Printf("error:%v\n", err)
		return
	}
	val := rdb.Get("testKey")
	fmt.Println(val.Val())
}
