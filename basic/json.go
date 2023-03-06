package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct{
	Title string `json:"title"`  //需要大写才可被json调用
	Year int     `json:"year"`
	Price int	 `json:"rmb"`
	Actors []string  `json:"actors"`
}


func main(){
	movie := Movie{"喜剧之王",2000,10,[]string{"周星驰","张柏芝"}}

	//编码的过程 结构体--->json
	jsonStr, err :=	json.Marshal(movie)
	if err != nil{
		fmt.Println("json marshal error!",err)
		return
	}
	fmt.Printf("jsonStr = %s\n",jsonStr)

	//解码的过程  jsonstr--->结构体
	//jsonStr =  {"title":"喜剧之王","year":2000,"rmb":10,"actors":["周星驰","张柏芝"]}
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil{
		fmt.Println("json unmarshal error!",err)
		return
	}
	fmt.Printf("myMovie = %v\n",my_movie)

}

