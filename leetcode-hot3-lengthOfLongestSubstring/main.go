package main

func main() {
	println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	maxLen := 0
	l, r := 0, 0

	for r < len(s) {
		_, ok := m[s[r]]
		if !ok {
			m[s[r]] = 1
			r++
		} else {
			delete(m, s[l])
			l++
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
