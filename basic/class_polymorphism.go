package main

import "fmt"

//本质是一个指针
type AnimalIF interface{
	Sleep() 
	GetColor() string //获取动物的颜色
	GetType() string //获取动物的种类

}

//具体的猫
type Cat struct{
	color string  //猫的颜色
}

func (this *Cat) Sleep(){
	fmt.Println("Cat is Sleep")
}

func (this *Cat) GetColor() string{
	return this.color
}

func (this *Cat) GetType() string{
	return "Cat"
}


//具体的类
type Dog struct{
	color string  //狗的颜色
}

func (this *Dog) Sleep(){
	fmt.Println("Dog is Sleep")
}

func (this *Dog) GetColor() string{
	return this.color
}

func (this *Dog) GetType() string{
	return "Dog"
}

func showAnimal(animal AnimalIF){
	animal.Sleep()
	fmt.Println("color = ",animal.GetColor())
	fmt.Println("kind = ",animal.GetType())
}

func main(){
	var animal AnimalIF
	animal = &Cat{"Green"}
	animal.Sleep() //调用的就是Cat的Sleep()方法，多态的现象

	animal = &Dog{"Yellow"}
	animal.Sleep()

	cat := Cat{"Green"}
	dog := Dog{"Yellow"}

	showAnimal(&cat)
	showAnimal(&dog)
}
