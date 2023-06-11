package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

//⾼效的轻量级第三⽅多路复⽤器——HttpRouter。

//第三个参数Params 包含了具名参数，具名参数的值可以在处理器内部通过ByName⽅法获取
//地址为 http://127.0.0.1:8080/hello/foo，则展示hello, foo!
func hello3(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
func main() {
	//调⽤New函数来创建⼀个多路复⽤器
	mux := httprouter.New()
	//这个程序不再使⽤HandleFunc 绑定处理器函数，⽽是直接把处理器函数与给定的HTTP⽅法进⾏绑定
	mux.GET("/hello/:name", hello3)
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
