package main

import "fmt"

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// 扩容到原来的两倍
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		//表示创建一个slice，length是zlen，容量是zcap。容量不能比长度小
		z = make([]int, zlen, zcap)
		// a built-in function; see text，把x复制到z上面
		copy(z, x)
	}
	z[len(x)] = y
	return z
}
