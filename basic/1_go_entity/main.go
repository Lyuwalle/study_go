package main

import (
	"flag"
	"fmt"
)

func main() {
	var name string
	flag.StringVar(&name, "name1", "everyone", "The greeting object.")
	flag.Parse()
	fmt.Printf("Hello, %v!\n", name)
}
