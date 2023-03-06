package lib1

import "fmt"

//当前lib1包提供的API
func Lib1Test(){  //函数首字母大写，函数对外开放，否则只能在当前包内调用
	fmt.Println("Lib1Test() ...")
}

func init(){
	fmt.Println("lib1 init() ...")
}