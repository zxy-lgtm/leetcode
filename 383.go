package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	map_m := make(map[rune]int)
	for _, v := range magazine {
		map_m[v]++
	}
	for _, v := range ransomNote {
		map_m[v]--
	}
	for _, v := range map_m {
		if v < 0 {
			return false
		}
	}
	return true
}
