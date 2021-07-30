package leetcode

import "sort"

//哈希表
func intersect(nums1 []int, nums2 []int) []int {
	len1, len2 := len(nums1), len(nums2)
	if len1 > len2 {
		return intersect(nums2, nums1)
	}
	ans := make([]int, 0, len1)
	m := map[int]int{}
	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if m[num] > 0 {
			ans = append(ans, num)
			m[num]--
		}
	}
	return ans
}

//排序后双指针
func intersect2(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)
	length1, length2 := len(nums1), len(nums2)
	index1, index2 := 0, 0

	intersection := []int{}
	for index1 < length1 && index2 < length2 {
		if nums1[index1] < nums2[index2] {
			index1++
		} else if nums1[index1] > nums2[index2] {
			index2++
		} else {
			intersection = append(intersection, nums1[index1])
			index1++
			index2++
		}
	}
	return intersection
}
