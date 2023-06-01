package main
/*
这是一个go web应用程序
启动方式：命令行cd到这个程序所在的目录，运行go run server.go
本地浏览器输入localhost:8080/***
*/
import (
	"fmt"
	"net/http"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	//把之前定义的handler 函数设置成根（root）URL（/ ）被访问时的处理器，然后启动服务器 并让它监听系统的8080端⼝
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
