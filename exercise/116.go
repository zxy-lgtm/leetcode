package leetcode

//Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	que := []*Node{root}

	for len(que) > 0 {
		tmp := que
		que = nil

		for i, node := range tmp {
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			if node.Left != nil {
				que = append(que, node.Left)
			}

			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}

	return root
}

func connect2(root *Node) *Node {
	if root == nil {
		return root
	}

	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		for node := leftmost; node != nil; node = node.Next {
			node.Left.Next = node.Right

			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}

func connect3(root *Node) *Node {
	if root == nil {
		return root
	}
	connecthelp(root.Left, root.Right)

	return root
}

func connecthelp(node1, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	connecthelp(node1.Left, node1.Right)
	connecthelp(node1.Right, node2.Left)
	connecthelp(node2.Left, node2.Right)

	return
}
