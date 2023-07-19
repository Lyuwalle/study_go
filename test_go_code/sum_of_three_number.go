package main

import (
	"fmt"
	"sort"
	"strconv"
)

//给定数组，计算三数之和等于零的所有数组
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return ans
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := n - 1
		for left < right {
			temp := nums[i] + nums[left] + nums[right]
			if temp == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				for left < right && nums[left] == nums[left+1] {
					left++
				}
			} else if temp < 0 {
				left++
			} else {
				right--
			}
		}
	}
	return ans
}

func main() {
	//将int类型的参数转成10进制
	i1 := strconv.FormatInt(1689238826565, 10)
	fmt.Println(i1)
}
