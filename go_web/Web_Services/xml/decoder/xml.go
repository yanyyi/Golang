package main

import (
	"encoding/xml"
	"fmt"
	"os"
)

// 定义一些结构,用于表示数据
type Post struct {
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func main() {
	xmlFile, err := os.Open("Web_Services/xml/encoder/post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	//xmlData, err := ioutil.ReadAll(xmlFile)
	//if err != nil {
	//	fmt.Println("Error reading XML data:", err)
	//	return
	//}
	//var post Post
	//xml.Unmarshal(xmlData, &post)
	decoder := xml.NewDecoder(xmlFile) //根据给定的XML数据生成相应的解码器
	for {

	}

	fmt.Println(post)
	fmt.Println("\n")
	fmt.Println(post.Xml)

}
