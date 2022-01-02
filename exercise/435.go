package exercise

import "sort"

type sortable [][]int

func (s sortable) Len() int {
	return len(s)
}

func (s sortable) Less(i, j int) bool {
	return s[i][1] < s[j][1]
}

func (s sortable) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func eraseOverlapIntervals(intervals [][]int) int {
	//sort.Sort(sortable(intervals))
	sort.Slice(intervals, func(i int, j int) bool { return intervals[i][1] < intervals[j][1] })
	ans, end := 1, intervals[0][1]

	for _, nums := range intervals[1:] {
		if nums[0] >= end {
			ans++
			end = nums[1]
		}
	}
	return len(intervals) - ans
}
