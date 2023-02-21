package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈的使用
type Stack struct {
	MaxTop int    //表示我们栈最大可以存放数个数
	Top    int    //表示栈顶,因为栈顶固定,因此我们直接使用Top
	arr    [5]int //数组模拟栈
}

func (this *Stack) Push(val int) (err error) {
	//先判断栈是否满了
	if this.Top == this.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	this.Top++
	//放入数据
	this.arr[this.Top] = val
	return
}

// 遍历栈,注意需要从栈顶开始遍历
func (this *Stack) List() {

}

func main() {
	fmt.Println("hello stack")
}
