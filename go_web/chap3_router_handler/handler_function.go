package main

import (
	"fmt"
	"net/http"
)

//处理器函数实际上就是与处理器拥有相同⾏为的函数：这些函数与 ServeHTTP ⽅法拥有相同的签名，也就是说，它们接受 ResponseWriter
//和指向Request 结构的指针作为参数
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}
func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/world", world)
	server.ListenAndServe()
}
