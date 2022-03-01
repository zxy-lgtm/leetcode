package exercise

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	left := head
	var isPalindromeHelp func(head *ListNode) bool
	isPalindromeHelp = func(head *ListNode) bool {
		if head == nil {
			return true
		}
		ok := isPalindromeHelp(head.Next)
		ok = ok && (left.Val == head.Val)
		left = left.Next
		return ok
	}
	return isPalindromeHelp(head)
}

// 比上面那个快
func isPalindrome_(head *ListNode) bool {
	tmp := head
	nums := make([]int, 0)
	for tmp != nil {
		nums = append(nums, tmp.Val)
		tmp = tmp.Next
	}
	for i := 0; i < len(nums)/2; i++ {
		if nums[i] != nums[len(nums)-1-i] {
			return false
		}
	}

	return true
}
