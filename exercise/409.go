package exercise

func longestPalindrome(s string) (l int) {
	m := make(map[rune]int)
	for _, k := range s {
		m[k]++
	}
	tag := 0
	for _, k := range m {
		l += (k / 2) * 2
		if k%2 == 1 {
			tag = 1
		}
	}
	return l + tag
}
