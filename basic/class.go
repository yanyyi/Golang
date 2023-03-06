package main

import "fmt"

//如果类名首字母大写，表示其他包也能够访问
type Hero struct{
	//如果说类的属性首字母大写，表示改属性对外能够访问，否则的话只能够类内访问
	Name string
	Ad int 
	level int  //level首字母小写，私有
}

func (this *Hero) Show(){
	fmt.Println("Name = ",this.Name)
	fmt.Println("Ad = ",this.Ad)
	fmt.Println("Level = ",this.level)
}

func (this *Hero) GetName() string{
	return this.Name
}
func (this *Hero) SetName(newName string){
	this.Name = newName
}


func main(){
	//创建一个对象
	hero := Hero{Name:"Zhangsan",Ad:100,level:1}
	hero.Show()
	hero.SetName("Mike")
	hero.Show()
	var heroname string
	heroname = hero.GetName()
	fmt.Println("hero name = ",heroname)

}
