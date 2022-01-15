package exercise

func findRepeatedDnaSequences(s string) (res []string) {
	m := make(map[string]int)
	for i := 0; i <= len(s)-10; i++ {
		m[s[i:i+10]]++
	}

	for i, k := range m {
		if k > 1 {
			res = append(res, i)
		}
	}

	return
}
