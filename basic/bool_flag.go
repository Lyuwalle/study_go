package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep1 = flag.String("s", "sss", "separator")
func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep1))
	if !*n {
		fmt.Println()
	}
}
