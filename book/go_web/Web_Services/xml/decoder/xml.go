package main

import (
	"encoding/xml"
	"fmt"
	"io"
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
	xmlFile, err := os.Open("Web_Services/xml/decoder/post.xml")
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
		t, err := decoder.Token() //每进行一次迭代,就从解码器里面获取一个token
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding XML into tokens:", err)
			break
		}

		switch se := t.(type) { //检查token类型
		case xml.StartElement:
			if se.Name.Local == "comment" {
				var comment Comment
				decoder.DecodeElement(&comment, &se) //将xml数据解码至结构体
				fmt.Println(comment)
			}
		}
	}

}
