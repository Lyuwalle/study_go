package main

import "fmt"

//匿名函数表示func关键字后没有函数名

//返回一个匿名函数，匿名函数的类型是int
func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}