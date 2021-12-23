package exercise

func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	//建立邻接图和入度统计
	graph := make(map[int][]int, n)
	degree := make([]int, n)
	for i := range edges {
		u, v := edges[i][0], edges[i][1]
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
		degree[v]++
		degree[u]++
	}

	//将度为1的全部加入队列
	queue := []int{}
	for i := 0; i < n; i++ {
		if degree[i] == 1 {
			queue = append(queue, i)
		}
	}

	cnt := len(queue)
	for n > 2 { //当节点数小于等于2的时候停止，这里不理解可以画图
		n -= cnt      //n去掉所有队列中的节点(这些节点度为1)
		for cnt > 0 { //遍历队列

			tmp := queue[0]
			queue = queue[1:] //queue.pop()

			degree[tmp] = 0                        //删去当前节点
			for i := 0; i < len(graph[tmp]); i++ { //遍历当前节点的邻接节点
				if degree[graph[tmp][i]] != 0 {
					degree[graph[tmp][i]]--         //去掉与当前节点的关系
					if degree[graph[tmp][i]] == 1 { //如果度为1了 就加入队列
						queue = append(queue, graph[tmp][i])
					}
				}
			}
			cnt--
		}
		cnt = len(queue)
	}
	ans := []int{}
	for _, v := range queue {
		ans = append(ans, v)
	}
	return ans
}
