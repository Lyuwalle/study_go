package main
//串联处理器和处理器函数
import (
	"fmt"
	"net/http"
	"reflect"
	"runtime"
)

//HandlerFunc 类型的函数
func hello2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

//log 函数接受⼀个HandlerFunc 类型的函数作为参数，然后返回另⼀个HandlerFunc 类型的函数作为值
//log 函数的返回值是⼀个匿名函数，因为这个匿名函数接受⼀个 ResponseWriter 和⼀个Request 指针作为参数，所以它实际上也
//是⼀个HandlerFunc 。在匿名函数内部，程序⾸先会获取被传⼊的HandlerFunc 的名字，然后再调⽤这个HandlerFunc 。
func log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
		fmt.Println("Handler function called - " + name)
		h(w, r)
	}
}
func main() {
	server := http.Server{Addr: "127.0.0.1:8080"}
	http.HandleFunc("/hello", log(hello2))
	server.ListenAndServe()
}
