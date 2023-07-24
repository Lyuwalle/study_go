package main

import (
	"fmt"
	"os"
)

func main() {
	var s,sep string
	//os.Args表示命令行参数 如go run temp2.go arg1 arg2 arg3
	for i := 1; i < len(os.Args); i++{
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}
