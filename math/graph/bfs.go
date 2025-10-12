package graph

// BFS performs a breadth-first search on the graph starting from the start node
// and returns the nodes in the order they were visited
func BFS(start Node) []Node {
	visited := make(map[Node]bool)
	queue := []Node{start}
	result := []Node{}

	visited[start] = true

	for len(queue) > 0 {
		// Dequeue the first node
		current := queue[0]
		queue = queue[1:]

		// Add to result
		result = append(result, current)

		// Visit all adjacent nodes
		postNodes := current.GetPostNodes()
		for _, node := range postNodes {
			if !visited[node] {
				visited[node] = true
				queue = append(queue, node)
			}
		}
	}

	return result
}
