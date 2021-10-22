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

**--非递归实现(迭代法)--**

### 先序遍历

```go
// import "container/list"
func preorderTraversal(root *TreeNode) (res []int) {

	if root == nil {
		return nil
	}

	var stack = list.New()

	stack.PushBack(root.Right)
	stack.PushBack(root.Left)
	res = append(res, root.Val)

	for stack.Len() > 0 {

		e := stack.Back()
		stack.Remove(e)
		node := e.Value.(*TreeNode) //e是Element类型，其值为e.Value.由于Value为接口，所以要断言

		if node == nil {
			continue
		}

		res = append(res, node.Val)
		stack.PushBack(node.Right)
		stack.PushBack(node.Left)

	}

	return res

}

```

### 后序遍历

```go
func postorderTraversal(root *TreeNode) (res []int) {

	if root == nil {
		return nil
	}

	var stack = list.New()

    stack.PushBack(root.Left)
    stack.PushBack(root.Right)
    res=append(res,root.Val)

    for stack.Len()>0 {

        e:=stack.Back()
        stack.Remove(e)
        node := e.Value.(*TreeNode)//e是Element类型，其值为e.Value.由于Value为接口，所以要断言
        
        if node==nil{
            continue
        }

        res=append(res,node.Val)
        stack.PushBack(node.Left)
        stack.PushBack(node.Right)

    }

    for i:=0;i<len(res)/2;i++{
        res[i],res[len(res)-i-1] = res[len(res)-i-1],res[i]
    }
    
    return res
}

```

### 中序遍历

```go
func inorderTraversal(root *TreeNode)(res []int) {

    if root==nil{
       return nil
    }

    stack:=list.New()
    node:=root
    //先将所有左节点找到，加入栈中
    for node!=nil{
        stack.PushBack(node)
        node=node.Left
    }
    //其次对栈中的每个节点先弹出加入到结果集中，再找到该节点的右节点的所有左节点加入栈中
    for stack.Len()>0{

        e:=stack.Back()
        node:=e.Value.(*TreeNode)
        stack.Remove(e)
        //找到该节点的右节点，再搜索他的所有左节点加入栈中
        res=append(res,node.Val)
        node=node.Right

        for node!=nil{
            stack.PushBack(node)
            node=node.Left
        }

    }
    
    return 
}

```
