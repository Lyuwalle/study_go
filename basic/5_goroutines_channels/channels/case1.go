//
package main

import "fmt"

func main() {
	//创建两个channel
	naturals := make(chan int)
	squares := make(chan int)

	// Counter
	// naturals通道接收x
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	// naturals通道把数据传给x，squares通道接收x * x
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}
	}()

	// Printer (in main goroutine)
	// 不断打印squares通道中的数据
	for {
		fmt.Println(<-squares)
	}
}
