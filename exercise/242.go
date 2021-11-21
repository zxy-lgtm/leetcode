package leetcode

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	map_s := make(map[rune]int)
	map_t := make(map[rune]int)
	for _, v := range s {
		map_s[v]++
	}
	for _, v := range t {
		map_t[v]++
	}
	for i, k := range map_s {
		if map_t[i] != k {
			return false
		}
	}
	return true
}
