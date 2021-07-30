package leetcode

import (
	"strings"
)

func checkInclusion(s1 string, s2 string) bool {
	head, tail := 0, 0
	len1, len2 := len(s1), len(s2)
	head = strings.LastIndex(s2, string(s1[0]))
	if head == -1 || len1 > len2 {
		return false
	}
	if head == len2-1 {
		head = head - len1 + 1
		tail = head + len1
	} else {
		tail = head + len1
		for tail > len2 {
			tail--
			head--
		}
	}

	var c1, c2 [26]int
	for _, ch := range s1 {
		c1[ch-'a']++
	}

	for head >= 0 {
		for _, ch := range s2[head:tail] {
			c2[ch-'a']++
		}
		if c1 == c2 {
			return true
		} else {
			head--
			tail--
			c2 = [26]int{0}
			if head < 0 {
				return false
			}
		}
	}
	return false
}
