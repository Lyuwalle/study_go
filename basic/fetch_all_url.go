package main

/*
使用并发goroutine请求所有url
所使用的时间不会超过执行时间最长的那个url
测试 https://golang.org http://gopl.io https://godoc.org 三个url
*/
import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	//创建了一个传递string类型参数的channel，只创建了一个channel
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		// start a goroutine, main本身表示一个goroutine，go表示创建一个新的goroutine
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}


func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	//io.Copy会把响应的内容copy到ioutil.Discard输出流中，
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}