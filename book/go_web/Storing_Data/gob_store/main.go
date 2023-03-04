package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io/ioutil"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func store(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	decoder := gob.NewDecoder(buffer)
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}

}

func main() {
	post := Post{Id: 1, Content: "Go语言从入门到放弃", Author: "Mike"}
	store(post, "post1")
	var postRead Post
	load(&postRead, "post1")
	fmt.Println(postRead)
}
