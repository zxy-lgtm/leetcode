package main

func isBadVersion(version int) bool

func firstBadVersion(n int) int {
	tail := n
	head := 1
	for head <= tail {
		mid := (head + tail) / 2
		if isBadVersion(mid) {
			tail = mid - 1
		} else {
			head = mid + 1
		}
	}
	return head
}
