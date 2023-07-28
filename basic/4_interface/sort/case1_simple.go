package main

import (
	"fmt"
	"sort"
)

//sort包内置的提供了根据一些排序函数来对任何序列排序的功能
//一个内置的排序算法需要知道三个东西：序列的长度，表示两个元素比较的结果，一种交换两个元素的方式
//package sort
//
//type Interface interface {
//    Len() int
//    Less(i, j int) bool // i, j are indices of sequence elements
//    Swap(i, j int)
//}

type StringSlice []string

//为了对序列进行排序，我们需要定义一个实现了这三个方法的类型，然后对这个类型的一个实例应用sort.Sort函数

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func main() {
	var sortedList = []string{"adc", "dau", "cbd"}
	sort.Sort(StringSlice(sortedList))
	fmt.Println(sortedList)
}