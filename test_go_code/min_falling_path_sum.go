package main

import "fmt"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	//创建一个长度（行数）为n的二维数组
	dp := make([][]int, n)
	for i := range dp {
		//创建长度为n的一维数组并赋值给dp[i]
		dp[i] = make([]int, n)
	}
	copy(dp[0], matrix[0])
	//除了第一行，其他行都按照贪心计算上一行累加的最小值，最终返回dp最后一行的最小的一个值就行了
	for i := 1; i < n; i++ {
		for j := 0; j < n; j++ {
			//
			mn := dp[i-1][j]
			if j > 0 {
				mn = min(mn, dp[i-1][j-1])
			}
			if j < n-1 {
				mn = min(mn, dp[i-1][j+1])
			}
			dp[i][j] = mn + matrix[i][j]
		}
	}
	minVal := dp[n-1][0]
	for _, val := range dp[n-1] {
		if val < minVal {
			minVal = val
		}
	}
	return minVal
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}


func main() {
	nums := []int{1,2,3}
	for _, val := range nums {
		fmt.Println(val)
	}
}
