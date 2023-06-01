package main

import (
	_"html/template"
	"net/http"
)

func main() {
	//创建了一个多路复用器
	mux := http.NewServeMux()
	//当服务器接收到⼀个以/static/ 开头的URL请求时，以下两⾏代码会移除URL中的/static/ 字符串，然后在public ⽬录中查找被请求的⽂件。
	//例如，当服务器接收到⼀个针对⽂件 http://localhost/static/css/bootstrap.min.css 的请求 时，它将会在public ⽬录中查找以下⽂件：<application root>/css/bootstrap.min.css
	files := http.FileServer(http.Dir("/public"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	//当有针对根URL的请求到达时，该请求就会被重定向到名 为index 的处理器函数
	mux.HandleFunc("/", index)
	//mux.HandleFunc("/err", err)
	//mux.HandleFunc("/login", login)
	//mux.HandleFunc("/logout", logout)
	//mux.HandleFunc("/signup", signup)
	//mux.HandleFunc("/signup_account", signupAccount)
	//mux.HandleFunc("/authenticate", authenticate)
	//mux.HandleFunc("/thread/new", newThread)
	//mux.HandleFunc("/thread/create", createThread)
	//mux.HandleFunc("/thread/post", postThread)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	//files := []string{"templates/layout.html",
	//	"templates/navbar.html",
	//	"templates/index.html"}
	//templates := template.Must(template.ParseFiles(files...))
	//threads, err := data.Threads()
	//if err == nil {
	//	templates.ExecuteTemplate(w, "layout", threads)
	//}
}
