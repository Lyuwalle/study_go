package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)
//运行方式：cd到本目录之下，
//go run postgres_sql.go
//前提，安装了postgres, 并且正确设置了环境变量（把安装包的bin文件夹下面的命令添加到环境变量里面），mac在.bash_profile里面修改
type Post_II struct {
	Id      int
	Content string
	Author  string
}

//定义一个指向sql.DB结构的指针
var Db *sql.DB

//go中的初始化函数
//init is called after all the variable declarations in the package have evaluated their initializers, and those are evaluated only after all the imported packages have been initialized.
func init() {
	var err error
	//要导入github.com/lib/pq，否则找不到postgres的driver
	//初始化Db，Open表示与数据库进行连接
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func Posts(limit int) (posts []Post_II, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post_II{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

func GetPost(id int) (post Post_II, err error) {
	post = Post_II{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

//方法名前面有一个指向Post结构的引用，表示方法的接收者（receiver），接收者可以不使⽤&符号，直接在⽅法内部对结构进⾏引⽤。
func (post *Post_II) create() (err error) {
	//预处理语句（sql语句模板）
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post_II) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2 and author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

func (post *Post_II) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

func main() {
	post := Post_II{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	post.create()
	fmt.Println(post)
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	posts, _ := Posts(1)
	fmt.Println(posts)

	readPost.Delete()
}




