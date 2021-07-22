package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	cnt := len(s)
	if cnt == 0 {
		return 0
	}

	tempMap := make(map[uint8]int)
	l, r, maxLen := 0, 1, 1
	tempMap[s[l]] = l
	for l < r && r < cnt {
		val, ok := tempMap[s[r]]
		if ok && val >= l {
			l = val + 1
		}

		if r-l+1 > maxLen {
			maxLen = r - l + 1
		}

		tempMap[s[r]] = r
		r++
	}

	return maxLen
}

func lengthOfLongestSubstring(s string) int {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i], string(s[i]))
		//fmt.Println(string(s[i]), "index:", index)
		if index == -1 {
			if i+1 > end {
				end = i + 1
				//fmt.Println("end", end)
			}
		} else {
			start += index + 1
			end += index + 1
			//fmt.Println("start:", start, "end:", end)
		}
	}
	return end - start
}

/*func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}*/
