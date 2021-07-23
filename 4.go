package main

/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	end1 := len(nums1)
	end2 := len(nums2)

	if end1 == 0 && end2 == 0 {
		return 0.00000
	}

	if end1 == 0 {
		if end2%2 == 1 {
			return float64(nums2[end2/2])
		}
		return float64(nums2[end2/2-1]+nums2[end2/2]) / 2.00000
	}

	if end2 == 0 {
		if end1%2 == 1 {
			return float64(nums1[end1/2])
		}
		return float64(nums1[end1/2-1]+nums1[end1/2]) / 2.00000
	}

	tip := 0

	start1 := 0
	start2 := 0
	fmt.Println(start1+start2 < (end1+end2)/2-1)

	for (start1 + start2) <= (end1+end2)/2-1 {
		fmt.Println("s")

		if nums1[start1] >= nums2[start2] && start2 < end2 {
			fmt.Println(start2)
			start2++
			tip = 2
		}

		if nums1[start1] < nums2[start2] && start1 < end1 {
			fmt.Println(start1)
			start1++
			tip = 1
		}
	}

	fmt.Println(end1 + end2)

	if (end1+end2)%2 == 1 {
		fmt.Println(4)
		fmt.Println(tip)
		if tip == 1 {
			fmt.Println(1)
			return float64(min(nums1[start1],nums2[start2]))
		}
		if tip == 2 {
			fmt.Println(2)
			return float64(min(nums1[start1],nums2[start2]))
		}
	}
	fmt.Println(3)
	return float64(nums1[start1]+nums2[start2]) / 2.00000
}*/

/*func main() {
	n1 := []int{1, 3}
	n2 := []int{2}
	findMedianSortedArrays(n1, n2)
}*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		midIndex := totalLength / 2
		return float64(getKthElement(nums1, nums2, midIndex+1))
	} else {
		midIndex1, midIndex2 := totalLength/2-1, totalLength/2
		return float64(getKthElement(nums1, nums2, midIndex1+1)+getKthElement(nums1, nums2, midIndex2+1)) / 2.0
	}
	return 0.00000
}

func getKthElement(nums1, nums2 []int, k int) int {
	index1, index2 := 0, 0
	for {

		//这里是有一个数组长度为零的情况
		if index1 == len(nums1) {
			return nums2[index2+k-1]
		}
		if index2 == len(nums2) {
			return nums1[index1+k-1]
		}

		//这里是两个数组长度都为1的情况和循环结束的情况
		if k == 1 {
			return min(nums1[index1], nums2[index2])
		}

		half := k / 2

		//这里是考虑k/2-1和数组本来长度的大小
		newIndex1 := min(index1+half, len(nums1)) - 1
		newIndex2 := min(index2+half, len(nums2)) - 1

		pivot1, pivot2 := nums1[newIndex1], nums2[newIndex2]

		if pivot1 <= pivot2 {
			k -= (newIndex1 - index1 + 1)
			index1 = newIndex1 + 1
		} else {
			k -= (newIndex2 - index2 + 1)
			index2 = newIndex2 + 1
		}
	}
	return 0.00000
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
