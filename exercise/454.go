package leetcode

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) (res int) {
	map_sum_12 := make(map[int]int)
	for _, k := range nums1 {
		for _, v := range nums2 {
			map_sum_12[k+v]++
		}
	}

	for _, k := range nums3 {
		for _, v := range nums4 {
			//if map_sum_12[0-k-v] != 0{
			res += map_sum_12[0-k-v]
			//}
		}
	}
	return
}
