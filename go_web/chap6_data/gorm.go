package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"time"
)

type PostIii struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	Comments  []Comment
	CreatedAt time.Time
}

//`sql:"not null"`表示该字段对应列的值不能为null 。
type Comment struct {
	Id        int
	Content   string
	Author    string `sql:"not null"`
	PostId    int    `sql:"index"`
	CreatedAt time.Time
}

var Db_i *gorm.DB

//Gorm可以通过⾃动数据 迁移特性来创建所需的数据库表，并在⽤户修改相应的结构时⾃动对 数据库表进⾏更新，
//所以这个程序⽆需使⽤setup.sql⽂件来设置数据库表：当我们运⾏这个程序时，程序所需的数据库表就会⾃动⽣成。
func init() {
	var err error
	Db_i, err = gorm.Open("postgres", "user=gwp dbname=gwp password=123456 sslmode=disable")
	if err != nil {
		panic(err)
	}
	Db_i.AutoMigrate(&PostIii{}, &Comment{})
}

func main() {
	post := PostIii{Content: "Hello World!", Author: "Sau Sheong"}
	fmt.Println(post)
	Db_i.Create(&post)
	fmt.Println(post)
	comment := Comment{Content: "Good post!", Author: "Joe"}
	Db_i.Model(&post).Association("Comments").Append(comment)
	var readPost PostIii
	Db_i.Where("author = $1", "Sau Sheong").First(&readPost)
	var comments []Comment
	Db_i.Model(&readPost).Related(&comments)
	fmt.Println(comments[0])
}
