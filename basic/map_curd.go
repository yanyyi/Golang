package main

import "fmt"

func printMap(city map[string]string){
	for key, value := range city{
		fmt.Println("key = ", key," value = ",value)
		
	}	
}

func changeMap(city map[string]string){  //引用传递，传递的是map指针
	city["England"] = "London"  //会修改传进来的map类型的原始值
}

func main(){
	cityMap := make(map[string]string)
	
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"

	//遍历
	for key, value := range cityMap{
		fmt.Println("key = ", key, " value = ", value)
		
	}

	//删除
	delete(cityMap, "China")

	//修改
	cityMap["USA"] = "Washington"


	//遍历
	fmt.Println("\n")
	printMap(cityMap)

	//修改
	changeMap(cityMap)
	fmt.Println("\n")
	printMap(cityMap)
}
