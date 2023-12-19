package main

import "fmt"

func main() {
	testCase1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(testCase1))
}

func maxArea(height []int) int {
	res := 0
	l := 0
	r := len(height) - 1

	for l < r {
		res = max(min(height[r], height[l])*(r-l), res)
		if l < r {
			l++
		} else {
			r--
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
