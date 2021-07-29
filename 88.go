package leetcode

func merge(nums1 []int, m int, nums2 []int, n int) {

	len := n + m

	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < n; i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	for {
		if m > 0 && n > 0 {
			if nums1[m-1] > nums2[n-1] {
				nums1[m-1], nums1[len-1] = 0, nums1[m-1]
				m--
			} else if nums1[m-1] <= nums2[n-1] {
				nums1[len-1] = nums2[n-1]
				n--
			}
		} else if m == 0 {
			nums1[len-1] = nums2[n-1]
			n--
		}
		len--
		if len <= 0 {
			return
		}
	}
}
