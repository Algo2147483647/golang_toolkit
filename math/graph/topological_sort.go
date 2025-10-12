package graph

func TopologicalSort(graph Graph) []Node {
	// 使用Kahn算法实现拓扑排序
	// 1. 计算每个节点的入度
	inDegree := make(map[Node]int)
	nodes := graph.GetNodes()
	edges := graph.GetEdges()

	// 初始化所有节点的入度为0
	for _, node := range nodes {
		inDegree[node] = 0
	}

	// 计算每个节点的实际入度
	for _, edge := range edges {
		from, _ := edge.GetNodes()
		inDegree[from]++
	}

	// 2. 将所有入度为0的节点加入队列
	queue := make([]Node, 0)
	for node, degree := range inDegree {
		if degree == 0 {
			queue = append(queue, node)
		}
	}

	// 3. 执行拓扑排序
	result := make([]Node, 0)
	for len(queue) > 0 {
		// 取出一个入度为0的节点
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		// 遍历该节点的所有后继节点
		postNodes := current.GetPostNodes()
		for _, postNode := range postNodes {
			// 将这些后继节点的入度减1
			inDegree[postNode]--
			// 如果某个后继节点的入度变为0，则将其加入队列
			if inDegree[postNode] == 0 {
				queue = append(queue, postNode)
			}
		}
	}

	// 检查是否存在环（如果结果中的节点数少于总节点数，说明存在环）
	if len(result) != len(nodes) {
		// 存在环，返回空切片表示无法进行拓扑排序
		return []Node{}
	}

	return result
}
