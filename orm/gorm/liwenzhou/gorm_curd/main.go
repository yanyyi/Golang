package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int64
}

func main() {
	//1.连接MySQL数据库
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//2.把模型与数据库中的表对应起来
	db.AutoMigrate(&User{})

	//3.创建
	//u := User{Name: "qimi", Age: 18} //在代码层面创建一个User对象
	////fmt.Println(db.NewRecord(&u))    //判断主键是否为空,相当于查看表中是否有数据 此处返回true
	//db.Create(&u)
	////fmt.Println(db.NewRecord(&u)) //  此处返回false
	//u2 := User{Name: "Stephen", Age: 25}
	//db.Create(&u2)

	//4.查询
	//user := new(User)
	//db.First(&user)
	//fmt.Printf("user: %s\n", *user)

	//var users []User
	//db.Find(&users)
	//fmt.Printf("users: %s\n", users)

	//5.更新
	//db.Debug().Model(&user).Where("name=?", "qimi").Update("name", "John")
	//db.Debug().Model(&user).Where("name=?", "John").Update(map[string]interface{}{"name": "Chen", "age": 50})
	//
	////只更新map中的某些字段
	//conitionMap := map[string]interface{}{
	//	"name": "Stephen",
	//	"age":  20,
	//}
	//db.Debug().Model(&user).Select("age").Update(conitionMap) //只更新age字段
	//db.Debug().Model(&user).Omit("age").Update(conitionMap)   //忽略age字段更新其他字段

	//6.删除
	db.Debug().Unscoped().Where("name=?", "Stephen").Delete(&User{})

}
