package exercise

import "sort"

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func merge(intervals [][]int) (res [][]int) {
	sort.Sort(sortable(intervals))
	tmp := 0
	for tmp < len(intervals) {
		res = append(res, intervals[tmp])
		for tmp < len(intervals) && res[len(res)-1][1] >= intervals[tmp][0] {
			tmp++
			if res[len(res)-1][1] < intervals[tmp-1][1] {
				res[len(res)-1][1] = intervals[tmp-1][1]
			}
		}

	}
	return
}
