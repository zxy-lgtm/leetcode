package exercise

func subarraySum(nums []int, k int) int {
	pre, count := 0, 0
	m := make(map[int]int)
	m[0] = 1
	for _, num := range nums {
		pre += num
		if _, ok := m[pre-k]; ok {
			count += m[pre-k]
		}
		m[pre]++
	}
	return count
}
