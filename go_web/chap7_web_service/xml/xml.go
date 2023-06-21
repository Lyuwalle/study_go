package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

//``表示结构标签（struct tag），对结构以及XML元素进⾏映射
//XMLName、类型为xml.Name的字段，可以将XML元素的名字存储在这个字段⾥⾯（在⼀般情况下，结构的名字就是元素的名字）。
//使用`xml:"id,attr"`可以将元素的id属性的值存储到这个字段⾥⾯。
//⽤`xml:",innerxml"`可以将XML元素中 的原始XML存储到这个字段⾥⾯。
type Post struct { //#A ❶
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
	xmlFile, err := os.Open("post.xml")
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}
	var post Post
	//将XML数据解封到结构⾥⾯
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}
