package main

import "fmt"

type Human struct {
	name string
	sex string
}

func (this *Human) Eat(){
	fmt.Println("Human Eat()...")
}

func (this *Human) Walk(){
	fmt.Println("Human Walk)...")
}

type SuperMan struct{
	Human //SuperMn类继承了Human类的方法
	level int
}

//重定义父类的方法
func (this *SuperMan) Eat(){
	fmt.Println("SuperMan Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly(){
	fmt.Println("SuperMan Fly()...")
}

func (this *SuperMan) Print(){
	fmt.Println("name = ",this.name)
	fmt.Println("sex = ",this.sex)
	fmt.Println("level = ",this.level)
}

func main(){
	h := Human{"Zhangsan","female"}
	h.Eat()
	h.Walk()
	fmt.Println("\n")
	//定义一个子类对象
	// s := SuperMan{Human{"Lisi","female"},88}
	var s SuperMan
	s.name = "John"
	s.sex = "man"
	s.level = 88

	s.Walk() //调用父类的方法
	s.Eat()  //调用子类的方法
	s.Fly() // 调用子类的方法
	s.Print()
}
