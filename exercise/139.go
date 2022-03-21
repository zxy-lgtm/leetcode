package exercise

func wordBreak(s string, wordDict []string) bool {
	mp := make(map[string]bool)

	for _, v := range wordDict {
		mp[v] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true

	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mp[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(s)]

}
