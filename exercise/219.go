package exercise

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, key := range nums {
		if j, ok := m[key]; ok && i-j <= k {
			return true
		}
		m[key] = i
	}
	return false
}
