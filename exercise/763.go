package exercise

func partitionLabels(s string) (res []int) {
	m := make(map[rune]int)
	for i, k := range s {
		m[k] = i
	}

	start, end := 0, 0
	for i, k := range s {
		if m[k] > end {
			end = m[k]
		}
		if i == end {
			res = append(res, end-start+1)
			start = end + 1
		}
	}
	return

}
