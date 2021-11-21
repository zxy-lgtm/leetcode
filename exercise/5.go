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
