package main

import (
	"fmt"
	"os"
)

func main() {
	//byte切片
	data := []byte("Hello World!\n")
	//向data1文件写入data
	err := os.WriteFile("data1", data, 0644)
	if err != nil {
		panic(err)
	}

	//读取data1文件并打印，返回值[]byte
	read1, _ := os.ReadFile("data1")
	fmt.Println(read1)

	//os创建文件data2，返回值*File，写入data
	file1, _ := os.Create("data2")
	defer file1.Close()
	bytes, _ := file1.Write(data)
	fmt.Printf("Wrote %d bytes to file\n", bytes)

	//os打开文件data2
	file2, _ := os.Open("data2")
	defer file2.Close()
	//创建一个byte切片
	read2 := make([]byte, len(data))
	bytes, _ = file2.Read(read2)
	fmt.Printf("Read %d bytes from file\n", bytes)
	fmt.Println(string(read2))
}
