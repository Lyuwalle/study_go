package main

import (
	"fmt"
	"net/http"
)

//给客户端设置cookie，运行后在浏览器可以看到
func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{Name: "first_cookie",Value: "go programming", HttpOnly: true}
	c2 := http.Cookie{Name: "second_cookie", Value: "hahaha", HttpOnly: true}
	c3 := http.Cookie{Name: "third_cookie", Value: "Yahaha", HttpOnly: true}
	w.Header().Set("Set-Cookie", c1.String())
	w.Header().Add("Set-Cookie", c2.String())
	//另外一种设置cookie的方式
	http.SetCookie(w, &c3)
}

//获取cookie，必须要先设置才能获取
func getCookie(w http.ResponseWriter, r *http.Request) {
	h := r.Header["Cookie"]
	fmt.Fprintln(w, h)
}

func main(){
	server := http.Server{
		Addr : "127.0.0.1:8080",
	}
	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)
	server.ListenAndServe()
}
