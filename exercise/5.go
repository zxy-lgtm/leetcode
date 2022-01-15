package main

import (
	"fmt"
)

func longestPalindrome(s string) string {
	dp := make([][]bool, len(s))
	fmt.Println(dp)
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
		//fmt.Println(dp[i])
	}
	ans := ""
	for i := 0; i < len(s); i++ {
		for k := 0; k <= i; k++ {
			//回文字符串它的最后一个字符等于第一个字符并且倒数第二个字符等于第二个字符并且不能越界
			dp[i][k] = s[i] == s[k] && (i-1 < k+1 || dp[i-1][k+1])

			//与上一个回文进行比较，找到最长
			if dp[i][k] && i-k+1 > len(ans) {
				ans = s[k : i+1]
				//fmt.Println(ans)
			}
		}
		//fmt.Println(dp[i])
	}
	//fmt.Println(dp[0], dp[1], dp[2], dp[3])
	return ans
}

/*func main() {

	fmt.Println(longestPalindrome("abba"))
}*/

func longestPalindrome_(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}
