package main

//将数组和减半的最少操作次数

import (
	"container/heap"
	"fmt"
)

// PriorityQueue 定义一个切片类型PriorityQueue 优先级队列，需要自己实现部分逻辑？
// 由heap包可知，要使用heap必须要实现五个方法：sort中的Len() int、Less(i, j int) bool、Swap(i, j int) heap中的Push(x any)、Pop() any
type PriorityQueue []float64

func (pq PriorityQueue) Len() int {
	return len(pq)
}

// Less 比较优先级队列中下标i，j中的值，因为每次都要弹出队列中的最大值，因此这里Less(i,j)的条件是pq[i] > pq[j]
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] > pq[j]
}

// Swap 交换pq中下标为i，j的值
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push 向pq中push一个数据
func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(float64))
}

// Pop 从pq中pop一个元素，pq元素个数减1
func (pq *PriorityQueue) Pop() any {
	old, n := *pq, len(*pq)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

//每次操作都选择当前数组的最大值进行减半操作。
func halveArray(nums []int) int {
	//定义一个指针，指向PriorityQueue类型的地址
	pq := &PriorityQueue{}
	sum, sum2 := 0.0, 0.0
	for _, num := range(nums) {
		heap.Push(pq, float64(num))
		sum += float64(num)
	}
	res := 0
	for sum2 < sum / 2 {
		x := heap.Pop(pq).(float64)
		sum2 += x / 2
		heap.Push(pq, x / 2)
		res ++
	}
	return res

}

func main() {
	fmt.Println(halveArray([]int{3,8,20}))
}
