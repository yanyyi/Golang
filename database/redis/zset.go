package main

import (
	"fmt"
	"redis/src"
)

func main() {
	err := src.initClient()
	if err != nil {
		fmt.Printf("init client failed, error:%v\n", err)
		return
	}
	fmt.Println("init connection success.")
	defer src.rdb.Close()

	zsetKey := "topn"
	//languages := []redis.Z{
	//	redis.Z{Score: 200, Member: "Java"},
	//	redis.Z{Score: 300, Member: "C++"},
	//	redis.Z{Score: 400, Member: "MySQL"},
	//	redis.Z{Score: 200, Member: "PHP"},
	//}
	//num, err := rdb.ZAdd(zsetKey, languages...).Result()
	//if err != nil {
	//	fmt.Printf("zadd failed, error:%v\n", err)
	//	return
	//}
	//fmt.Printf("zadd %d succ.\n", num)

	//newScore, err := rdb.ZIncrBy(zsetKey, 300, "PHP").Result()
	//if err != nil {
	//	fmt.Printf("zincrby failed, error:%v\n", err)
	//	return
	//}
	//fmt.Printf("PHP score is %d now", int(newScore))

	//num, err := rdb.ZRem(zsetKey, "MySQL").Result()
	//if err != nil {
	//	fmt.Printf("zrem failed, error:%v\n", err)
	//	return
	//}
	//fmt.Printf("integer: %d", int(num))

	//ret, err := rdb.ZRevRange(zsetKey, 0, -1).Result()
	//if err != nil {
	//	fmt.Printf("zrevrange failed, error:%v\n", err)
	//	return
	//}
	//fmt.Println(ret)  //返回一个列表[PHP C++ Java]

	ranking, err := src.rdb.ZRank(zsetKey, "C++").Result()
	if err != nil {
		fmt.Printf("zrank failed, error:%v\n", err)
		return
	}
	fmt.Printf("C++ ranking is %d", ranking)

}
