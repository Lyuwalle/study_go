package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var fileName = "posts.csv"

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//创建一个csv文件
	csvFile, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}
	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	//保证缓冲区中的所有数据都已经被正确地写⼊⽂件⾥⾯了
	writer.Flush()

	//打开csv文件
	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	//no check is made and records may have a variable number of fields.
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}
