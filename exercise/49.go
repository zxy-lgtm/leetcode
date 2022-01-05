package exercise

import "sort"

func groupAnagrams(strs []string) (res [][]string) {
	if len(strs) == 0 {
		return
	}

	m := make(map[string][]string)
	for _, k := range strs {
		s := []byte(k)
		sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
		m[string(s)] = append(m[string(s)], k)
	}
	for _, k := range m {
		res = append(res, k)
	}

	return

}
