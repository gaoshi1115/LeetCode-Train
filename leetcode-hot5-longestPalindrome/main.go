package main

import "fmt"

func main() {
	case1 := "a"
	fmt.Println(longestPalindrome(case1))
}

// 动态规划
func longestPalindrome(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := j; i >= 0; i-- {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && j-i+1 > len(res) {
				res = s[i : j+1]
			}
		}
	}
	return res
}
