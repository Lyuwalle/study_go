package main

import (
	"fmt"
	"sort"
)

//拓扑排序
//给定一些计算机课程，每个 课程都有前置课程，只有完成了前置课程才可以开始当前课程的学习；我们的目标是选择出 一组课程，
//这组课程必须确保按顺序学习时，能全部被完成。每个课程的前置课程如下：
var prereqs = map[string][]string{
	"algorithms":            {"data structures"},
	"calculus":              {"linear algebra"},
	"compilers":             {"data structures", "formal languages", "computer organization"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for i, class := range topo_sort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, class)
	}
}

//topo_sort 对参数prereqs进行拓扑排序
func topo_sort(prereqs map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	//对prereqs中的key进行排序
	var keys []string
	for item := range prereqs {
		keys = append(keys, item)
	}
	sort.Strings(keys)
	//用深度优先搜索了整张图，获得了符合要求的课程序列
	//当匿名函数需要被递归调用时，我们必须首先声明一个变量visitAll
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				//递归
				visitAll(prereqs[item])
				order = append(order, item)
			}
		}
	}
	visitAll(keys)
	return order
}
