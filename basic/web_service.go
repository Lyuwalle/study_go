package main
/*go run web_service.go & 命令让这个服务在后台跑起来*/
/*然后访问8000端口，则会显示所请求的路径*/
import (
	"fmt"
	"log"
	"net/http"
)

//main函数将所有发送到/路径下的请求和handler函数关联起来
func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}


// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
