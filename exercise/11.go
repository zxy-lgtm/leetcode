package exercise

func maxArea(height []int) int {
	max := 0
	start, end, lens := 0, len(height)-1, len(height)-1

	for start <= end {
		if height[start] < height[end] {
			if height[start]*lens > max {
				max = height[start] * lens
			}
			start++
		} else {
			if height[end]*lens > max {
				max = height[end] * lens
			}
			end--
		}
		lens--
	}

	return max
}
