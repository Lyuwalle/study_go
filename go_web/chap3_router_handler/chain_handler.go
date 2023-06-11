package main

import (
	"fmt"
	"net/http"
)

type HelloHandler2 struct{}

func (h HelloHandler2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func log2(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Handler called - %T\n", h)
		h.ServeHTTP(w, r)
	})
}
func protect(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//省略了⼀段⽤于检测⽤户登录情况的代码
		h.ServeHTTP(w, r)
	})
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	hello3 := HelloHandler2{}
	http.Handle("/hello", protect(log2(hello3)))
	server.ListenAndServe()
}
