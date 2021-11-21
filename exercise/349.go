package leetcode

func intersection(nums1 []int, nums2 []int) (res []int) {
	map_nums1 := make(map[int]int)
	for _, v := range nums1 {
		if map_nums1[v] == 0 {
			map_nums1[v]++
		}
	}
	for _, v := range nums2 {
		if map_nums1[v] != 0 {
			map_nums1[v]--
			res = append(res, v)
		}
	}
	return

}
