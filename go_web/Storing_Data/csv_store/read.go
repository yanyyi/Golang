package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post2 struct {
	Id      int
	Content string
	Author  string
}

func main() {

	file, err := os.Open("./posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post2
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post2{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
