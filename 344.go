package leetcode

func reverseString(s []byte) {
	head, tail := 0, len(s)-1
	for head < tail {
		s[head], s[tail] = s[tail], s[head]
		head++
		tail--
	}
	return

}
