package leetcode

func backspaceCompare(s string, t string) bool {
	return truestring(s) == truestring(t)
}

func truestring(s string) (res string) {
	b := []byte(s)
	r := make([]byte, 0)
	l := 0
	for _, v := range b {
		//fmt.Println(v)
		if v != '#' {
			r = append(r, byte(v))
			l++
		} else {
			if l != 0 {
				l--
				r = r[:l]
			}
		}
	}

	return string(r)
}
