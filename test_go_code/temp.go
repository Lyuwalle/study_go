package main

import (
	"fmt"
	"strings"
)

func add1(r rune) rune {
	return r + 1
}

func main() {
	//对HAL-9000每个字符都加1然后返回
	fmt.Println(strings.Map(add1, "HAL-9000"))
}


