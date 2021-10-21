package leetcode

//Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//利用一个k容量的切片按从小到大的顺序存储最小的k个元素，最后返回第k个元素
func kthSmallest(root *TreeNode, k int) int {

	min := make([]int, k)
	max := []int{root.Val}

	FMax(root, max)

	for i := k - 1; i >= 0; i-- {
		min[i] = max[0]
	}

	First(root, min)

	return min[k-1]

}

func FMax(node *TreeNode, max []int) {

	if node == nil {
		return
	}

	if max[0] < node.Val {
		max[0] = node.Val
	}

	FMax(node.Left, max)
	FMax(node.Right, max)

}

func First(node *TreeNode, min []int) {

	if node == nil {
		return
	}

	tag := node.Val

	for i, val := range min {
		if tag < val {
			temp := tag
			tag = val
			min[i] = temp
		}
	}
	First(node.Left, min)
	First(node.Right, min)

}
