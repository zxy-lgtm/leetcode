package leetcode

import "strings"

func reverseWords(s string) string {
	ss := strings.SplitN(s, " ", -1)
	var re func(string) string
	re = func(s string) string {
		sss := []byte(s)
		head, tail := 0, len(sss)-1
		for head < tail {
			sss[head], sss[tail] = sss[tail], sss[head]
			head++
			tail--
		}
		return string(sss)

	}
	for i, part := range ss {
		ss[i] = re(part)
	}

	ans := strings.Join(ss, " ")
	return ans

}
