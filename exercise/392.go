package exercise

func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	start := 0
	for _, ch := range t {
		if s[start] == byte(ch) {
			start++
		}
		if start == len(s) {
			return true
		}
	}
	return false
}
