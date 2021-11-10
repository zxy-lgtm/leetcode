package leetcode

import "strings"

func replaceSpace(s string) (res string) {
	strs := strings.Split(s, " ")
	var i int
	for i = 0; i < len(strs)-1; i++ {
		res += strs[i]
		res += "%20"
	}
	res += strs[i]
	return
}
