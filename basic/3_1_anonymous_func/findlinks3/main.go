package main

//广度优先算法抓取链接

import (
	"fmt"
	"log"
	"os"
	"study_go/basic/3_1_anonymous_func/links"
)

// breadthFirst calls f for each item in the worklist. 广度优先算法
// Any items returned by f are added to the worklist.
// f is called at most once for each item. f是一个函数类型参数，参数为string 返回值为[]string
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	//这里的循环条件:worklist的长度大于0，每个循环开始之前会把worklist置为nil
	//从一个url链接中拿到所有链接之后，在从这些链接中一个一个取子链接，直至所有的链接网全部被访问 程序运行结束
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				//append的参数“f(item)...”，会将f返回的一组元素一个个添加到worklist中。
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//从url中提取出所有链接，和links包下面的Extract函数一样
func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}



