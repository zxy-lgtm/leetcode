package exercise

func getMaxLen(nums []int) (ans int) {
	maxF, minF := 0, 0
	for _, num := range nums {
		if num > 0 {
			maxF++
			if minF > 0 {
				minF++
			}
		} else if num == 0 {
			maxF, minF = 0, 0
		} else {
			if minF > 0 {
				maxF, minF = minF+1, maxF+1
			} else {
				maxF, minF = 0, maxF+1
			}
		}
		ans = max(ans, maxF)
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getMaxLen_(nums []int) int {
	minF, maxF := 0, 0
	tmp := 0
	for _, v := range nums {
		if v > 0 {
			maxF++
			if minF > 0 {
				minF++
			}
		} else if v == 0 {
			maxF = 0
			minF = 0
		} else if v < 0 {
			if minF > 0 {
				maxF, minF = minF+1, maxF+1
			} else {
				minF, maxF = maxF+1, 0
			}
		}
		tmp = max(tmp, maxF)
	}
	return tmp
}
