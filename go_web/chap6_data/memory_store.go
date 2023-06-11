package main

import "fmt"
//内存存储
type Post struct {
	Id      int
	Content string
	Author  string
}

//为什么使用指针？
//两个map映射的都是指向帖子的指针，而不是帖子本身。这样保证通过id或者作者的名字获取到的帖子都是相同的帖子而不是同一个帖子的不同副本
//[]*Post表示切片，每个切片可以包含多个指向Post的指针

var PostById map[int]*Post
var PostByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = append(PostByAuthor[post.Author], &post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string][]*Post)

	post1 := Post{1, "Hello World!", "Foo"}
	post2 := Post{2, "Go Programming", "Google"}
	post3 := Post{3, "Go Web Programming", "Sau"}
	post4 := Post{4, "Go Web Programming2", "Sau"}


	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostByAuthor["Sau"] {
		fmt.Println(post)
	}
}