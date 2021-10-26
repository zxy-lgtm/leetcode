package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if subRoot == nil {
		return true
	}
	if root == nil && subRoot != nil {
		return false
	}

	ok := isSameTree(root, subRoot)
	if !ok {
		ok = isSubtree(root.Left, subRoot)
		if ok {
			return ok
		}
		ok = isSubtree(root.Right, subRoot)
		if ok {
			return ok
		}
	}

	return ok
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil {
		return false
	}
	if q == nil && p != nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	okl := isSameTree(p.Left, q.Left)
	okr := isSameTree(p.Right, q.Right)
	return okl && okr
}
