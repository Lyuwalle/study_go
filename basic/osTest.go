package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//内置函数make创建map，类似java的hashmap
	//map键是字符串，值是int
	//bufio处理输入和输出
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//不存在的键值设为0
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
