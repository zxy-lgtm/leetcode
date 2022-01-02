package exercise

func findAnagrams(s string, p string) (res []int) {
	l, lp := len(s), len(p)

	if l < lp {
		return
	}

	var S, P [26]int

	for i, k := range p {
		S[s[i]-'a']++
		P[k-'a']++
	}

	if S == P {
		res = append(res, 0)
	}

	for i, k := range s[:l-lp] {
		S[k-'a']--
		S[s[i+lp]-'a']++
		if S == P {
			res = append(res, i+1)
		}
	}
	return
}
