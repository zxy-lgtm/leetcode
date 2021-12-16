package exercise

func findOrder(numCourses int, prerequisites [][]int) []int {
	var (
		edges   = make([][]int, numCourses)
		visited = make([]int, numCourses) // 是否已经被检查，是的话为2
		result  []int
		valid   bool = true //是否有环
		dfs     func(u int)
	)

	dfs = func(u int) {
		visited[u] = 1
		for _, v := range edges[u] {
			// fmt.Println(v)
			if visited[v] == 0 {
				dfs(v)
				if !valid {
					return
				}
			} else if visited[v] == 1 { // 如果回到了待搜索完成的一个节点，那么说明这个必定有一个环，不满足条件
				valid = false
				return
			}
		}
		visited[u] = 2
		result = append(result, u) // 将完成搜索的节点加入到结果中
	}

	for _, info := range prerequisites {
		edges[info[1]] = append(edges[info[1]], info[0])
		// fmt.Println(edges)
	}

	for i := 0; i < numCourses && valid; i++ {
		if visited[i] == 0 {
			dfs(i)
		}
	}
	if !valid {
		return []int{}
	}
	for i := 0; i < len(result)/2; i++ {
		result[i], result[numCourses-i-1] = result[numCourses-i-1], result[i] // 倒序
	}
	return result
}
