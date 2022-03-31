package exercise

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	tl, sl := 0, 0
	for tl < len(t) {
		if s[sl] == t[tl] {
			sl++
		}
		tl++
		if sl == len(s) {
			return true
		}
	}

	return false
}
