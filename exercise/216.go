package exercise

var res [][]int

func combinationSum3(k int, n int) [][]int {
	res = [][]int{}
	if min_sum(k) > n {
		return res
	}
	backtrack(k, n, 1, []int{})
	return res

}

func backtrack(k, n, start int, track []int) {
	if len(track) == k && array_sum(track) == n {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
	}
	if array_sum(track) > n {
		return
	}
	for i := start; i < 10; i++ {
		track = append(track, i)
		backtrack(k, n, i+1, track)
		track = track[:len(track)-1]
	}
}

func min_sum(k int) (sum int) {
	for k > 0 {
		sum += k
		k--
	}
	return
}

func array_sum(track []int) (sum int) {
	for _, num := range track {
		sum += num
	}
	return
}
