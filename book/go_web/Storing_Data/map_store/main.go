package main

import "fmt"

type Book struct {
	Id      int
	Content string
	Author  string
}

var BookById map[int]*Book
var BookByAuthor map[string][]*Book

func store(b Book) {
	BookById[b.Id] = &b
	BookByAuthor[b.Author] = append(BookByAuthor[b.Author], &b)
}

func main() {
	BookById = make(map[int]*Book)
	BookByAuthor = make(map[string][]*Book)

	book1 := Book{Id: 1, Content: "三体", Author: "刘慈欣"}
	book2 := Book{Id: 2, Content: "黑暗森林", Author: "刘慈欣"}
	book3 := Book{Id: 3, Content: "死神永生", Author: "刘慈欣"}
	book4 := Book{Id: 4, Content: "机器学习", Author: "周志华"}
	book5 := Book{Id: 5, Content: "信号与系统", Author: "Oppenheim"}
	store(book1)
	store(book2)
	store(book3)
	store(book4)
	store(book5)

	fmt.Println(*BookById[5])

	for _, book := range BookByAuthor["刘慈欣"] {
		fmt.Println(*book)
	}

}
