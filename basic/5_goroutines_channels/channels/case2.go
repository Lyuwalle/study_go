package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// naturals通道接收100个数据之后就关闭通道
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		// 当一个channel被关闭后，再向该channel发送数据将导致panic异常
		close(naturals)
	}()

	// Squarer
	// 从naturals通道中拿到所有数据之后，写入squares通道，然后关闭squares通道
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// Printer (in main goroutine)
	for x := range squares {
		fmt.Println(x)
	}
}

