package exercise

func trap(height []int) int {
	maxL, maxR := 0, 0
	l, r := 0, len(height)-1
	all := 0
	for l < r {
		maxL = max(height[l], maxL)
		maxR = max(height[r], maxR)
		if maxL < maxR {
			all += maxL - height[l]
			l++
		} else {
			all += maxR - height[r]
			r--
		}
	}
	return all
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
