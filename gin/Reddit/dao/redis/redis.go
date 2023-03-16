package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// 声明一个全局的rdb变量
var client *redis.Client

// 初始化连接
func Init() (err error) {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		Password: viper.GetString("redis.password"), // no password set
		DB:       viper.GetInt("redis.db"),          // use default DB
		PoolSize: viper.GetInt("redis.pool_size"),
	})

	_, err = client.Ping().Result()
	return err
}

func Close() {
	_ = client.Close()
}
