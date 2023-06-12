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

//存储文件, 将帖⼦存储为⼆进制数据
//参数
//data：空接口
//filename 被存储的二进制文件的名字
func store2(data interface{}, filename string) {
	//可变长度字节缓冲区，bytes.Buffer 既是读取器也是写⼊器
	buffer := new(bytes.Buffer)
	//把缓冲区传递给NewEncoder函数，来创建一个gob编码器
	encoder := gob.NewEncoder(buffer)
	//将数据编码到缓冲区里面
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	//缓冲区里面的已编码的数据写入文件
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

//加载文件, 读取这些⼆进制数据来获取帖⼦
func load2(data interface{}, filename string) {
	//读取二进制文件
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	//放入缓冲区
	buffer := bytes.NewBuffer(file)
	decoder := gob.NewDecoder(buffer)
	//Decode reads the next value from the input stream and
	//stores it in the data represented by the empty interface value.
	err = decoder.Decode(data)
	if err != nil {
		panic(err)
	}
}

func main() {
	filename := "post1"
	post := Post{1, "Hello World!", "Sau"}
	store2(post, filename)
	var postRead Post
	//注意这里是传指向postRead的指针，向里面放数据
	load2(&postRead, filename)
	fmt.Println(postRead)
}
