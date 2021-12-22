package exercise

//Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

//深度优先
func cloneGraph(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := map[*Node]*Node{}
	var cg func(node *Node) *Node
	cg = func(node *Node) *Node {
		if _, ok := visited[node]; ok {
			return visited[node]
		}

		c := &Node{node.Val, []*Node{}}
		visited[node] = c

		for _, neighbor := range node.Neighbors {
			c.Neighbors = append(c.Neighbors, cg(neighbor))
		}

		return c
	}

	return cg(node)

}

//广度优先
func cloneGraph_(node *Node) *Node {
	if node == nil {
		return node
	}

	visited := map[*Node]*Node{}
	queue := []*Node{node}

	visited[node] = &Node{node.Val, []*Node{}}

	for len(queue) != 0 {
		c := queue[0]
		queue = queue[1:]
		for _, neighbor := range c.Neighbors {
			if _, ok := visited[neighbor]; !ok {
				cl := &Node{neighbor.Val, []*Node{}}
				visited[neighbor] = cl
				queue = append(queue, neighbor)
			}
			visited[c].Neighbors = append(visited[c].Neighbors, visited[neighbor])
		}
	}
	return visited[node]

}
