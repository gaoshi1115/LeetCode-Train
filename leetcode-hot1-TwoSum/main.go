package main

import (
	"fmt"
)

func main() {
	a := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(a)
}

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for i, v := range nums {
		if _, ok := indexMap[target-v]; ok {
			return []int{i, indexMap[target-v]}
		}
		indexMap[v] = i
	}
	return nil
}
