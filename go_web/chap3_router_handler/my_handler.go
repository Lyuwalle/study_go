package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct{}

//任何接⼝只要拥有⼀个 ServeHTTP ⽅法, 并且有ServeHTTP(http.ResponseWriter, *http.Request)签名，那么它就是一个处理器
func (h *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
func main() {
	handler := MyHandler{}
	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
