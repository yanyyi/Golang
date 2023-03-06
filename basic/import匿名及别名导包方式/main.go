package main 

import (
	_"goprojects/chapter03/lib1" //匿名导包，使得可以继续调用lib1中的init方法
	mylib2 "goprojects/chapter03/lib2" //别名导包
)

 
func main(){
	// lib1.Lib1Test()
	mylib2.Lib2Test()
}