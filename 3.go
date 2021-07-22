package main

import (
	"strings"
)

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
