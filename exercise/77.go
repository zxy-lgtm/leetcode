package exercise

var res [][]int

func combine(n int, k int) [][]int {
	res = [][]int{}
	backtrack(n, k, 1, []int{})
	return res
}

func backtrack(n, k, start int, track []int) {
	if len(track) == k {
		temp := make([]int, k)
		copy(temp, track)
		res = append(res, temp)
	}
	if len(track)+n-start+1 < k {
		return
	}
	for i := start; i <= n; i++ {
		track = append(track, i)
		backtrack(n, k, i+1, track)
		track = track[:len(track)-1]
	}
}
