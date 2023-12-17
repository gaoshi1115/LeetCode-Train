package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("ABCdf"))
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	maxLen := 0
	l, r := 0, 0
	for r < len(s) {
		tmp, ok := m[s[r]]
		if !ok {
			m[s[r]] = r
			r++
		} else {
			delete(m, s[l])
			l = tmp + 1
		}
		maxLen = max(maxLen, r-l)
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
