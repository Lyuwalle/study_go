package main

import "fmt"

//输入：jewels = "aA", stones = "aAAbbbb"
//输出：3
func numJewelsInStones(jewels string, stones string) int {
	count := 0
	//字符串转化成切片数组
	stoneSlice := []rune(stones)
	jewelSlice := []rune(jewels)
	for _, stone := range stoneSlice {
		for _, jewel := range jewelSlice {
			if stone == jewel {
				count ++
			}
		}
	}
	return count
}

//字符串可以直接以字符下标的方式访问
func numJewelsInStones2(jewels string, stones string) int {
	jewelsCount := 0
	for _, s := range stones {
		for _, j := range jewels {
			if s == j {
				jewelsCount++
				break
			}
		}
	}
	return jewelsCount
}

//哈希表方法
func numJewelsInStones3(jewels string, stones string) int {
	jewelsCount := 0
	jewelsSet := map[byte]bool{}
	for i := range jewels {
		jewelsSet[jewels[i]] = true
	}
	for i := range stones {
		if jewelsSet[stones[i]] {
			jewelsCount++
		}
	}
	return jewelsCount
}

func main() {
	fmt.Println(numJewelsInStones("aA", "aABA"))
}
