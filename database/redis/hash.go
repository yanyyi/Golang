package main

//func main() {
//	if err := initClient(); err != nil {
//		fmt.Printf("init client failed, err: %v\n", err)
//	}
//	fmt.Println("connect redis success...")
//	defer rdb.Close()

//rdbSet := rdb.HSet("user:1001", "id", 1)
//if rdbSet.Val() {
//	fmt.Println("设置成功")
//} else {
//	fmt.Println("设置失败")
//}

//rdbSet := rdb.HSet("user:1001", "name", "Mike")
//if rdbSet.Val() {
//	fmt.Println("设置成功")
//} else {
//	fmt.Println("设置失败")
//}

//val, err := rdb.HGet("user:1001", "name").Result()
//if err != nil {
//	fmt.Printf("error:%v\n", err)
//	return
//} else {
//	fmt.Println(val)
//}

//exists, err := rdb.HExists("user:1001", "name").Result()
//if err != nil {
//	fmt.Printf("error:%v\n", err)
//	return
//} else {
//	fmt.Println("exists: ", exists)
//}

//fmt.Println(rdb.HKeys("user:1001").Val())
//}
