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

//优化
func checkInclusion2(s1, s2 string) bool {
	n, m := len(s1), len(s2)
	if n > m {
		return false
	}
	cnt := [26]int{}
	for i, ch := range s1 {
		cnt[ch-'a']--
		cnt[s2[i]-'a']++
	}
	diff := 0
	for _, c := range cnt[:] {
		if c != 0 {
			diff++
		}
	}
	if diff == 0 {
		return true
	}
	for i := n; i < m; i++ {
		x, y := s2[i]-'a', s2[i-n]-'a'
		if x == y {
			continue
		}
		if cnt[x] == 0 {
			diff++
		}
		cnt[x]++
		if cnt[x] == 0 {
			diff--
		}
		if cnt[y] == 0 {
			diff++
		}
		cnt[y]--
		if cnt[y] == 0 {
			diff--
		}
		if diff == 0 {
			return true
		}
	}
	return false
}
