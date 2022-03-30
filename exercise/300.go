package exercise

func lengthOfLIS(nums []int) int {
	maxL := 0
	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
		if dp[i] > maxL {
			maxL = dp[i]
		}
	}
	return maxL
}
