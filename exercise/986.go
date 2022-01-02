package exercise

func intervalIntersection(firstList [][]int, secondList [][]int) (res [][]int) {
	len1, len2 := len(firstList), len(secondList)
	if len1 == 0 || len2 == 0 {
		return
	}

	for len(firstList) > 0 && len(secondList) > 0 {
		f, s := firstList[0], secondList[0]
		ans := compare(f, s)

		if ans != nil {
			res = append(res, ans)
		}

		if f[1] < s[1] {
			firstList = firstList[1:]
		} else {
			secondList = secondList[1:]
		}
	}
	return res
}

func compare(f, s []int) []int {
	start := max(f[0], s[0])
	end := min(f[1], s[1])
	if start > end {
		return nil
	}

	return []int{start, end}
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}

	return i
}
