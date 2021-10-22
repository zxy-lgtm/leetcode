## 二叉树的三种深度优先遍历

**--递归实现 --**

### 前序遍历

```go
func preorderTraversal(root *TreeNode) (res []int) {
    
	var traversal func(node *TreeNode)
    
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
        
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
        
	}
	traversal(root)
	return res
}

```



### 中序遍历

```go
func inorderTraversal(root *TreeNode) (res []int) {
    
	var traversal func(node *TreeNode)
    
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
        
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
        
	}
	traversal(root)
	return res
}

```



### 后序遍历

```go
func postorderTraversal(root *TreeNode) (res []int) {

	var traversal func(node *TreeNode)

	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}

		traversal(node.Left)
		traversal(node.Right)
		res = append(res, node.Val)

	}
	traversal(root)
	return res
}

```
