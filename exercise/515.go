package leetcode

import (
	"container/list"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) (finRes []int) {

	res := [][]int{}

	if root == nil {
		return
	}

	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int

	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode) //出队列
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr)
		tmpArr = []int{}
	}
	// 找到每层的最大值
	for i := 0; i < len(res); i++ {
		finRes = append(finRes, max(res[i]...))
	}
	return
}
func max(vals ...int) int {

	max := int(math.Inf(-1)) //负无穷

	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
