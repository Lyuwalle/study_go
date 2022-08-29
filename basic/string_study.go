package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "hello, world"
	s1 := "hello, 世界"
	s2 := "hello, せかい"
	fmt.Println(len(s))            // "12"
	fmt.Println(s[0], s[7])        // "104 119" ('h' and 'w')表示第0个和第7个字节的字节值
	fmt.Println(s[0:5])            // "hello"
	fmt.Println(s[:5])             // "hello"
	fmt.Println(s[7:])             // "world"
	fmt.Println(s[:])              // "hello, world"
	fmt.Println("goodbye" + s[5:]) // "goodbye, world"

	fmt.Println(len(s1))                    //13
	fmt.Println(utf8.RuneCountInString(s1)) //9
	fmt.Println(len(s2))                    //16
	fmt.Println(utf8.RuneCountInString(s2)) //10

}

// HasPrefix 字符串s的前缀是否是prefix
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

// Contains 字符串s中是否有substr子串
func Contains(s, substr string) bool {
	for i := 0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}
