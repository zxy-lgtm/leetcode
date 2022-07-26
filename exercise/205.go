package exercise

func isIsomorphic(s string, t string) bool {
	return isIsomorphic_h(s, t) && isIsomorphic_h(t, s)
}

func isIsomorphic_h(s string, t string) bool {
	m := make(map[byte]byte)
	for i, ch := range s {
		if m[byte(ch)] != 0 && (m[byte(ch)] != t[i]) {
			return false
		}
		m[byte(ch)] = t[i]
	}
	return true
}
