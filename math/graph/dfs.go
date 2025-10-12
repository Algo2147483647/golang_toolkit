package graph

// DFS performs a depth-first search on the graph starting from the start node
// and returns the nodes in the order they were visited
func DFS(start Node) []Node {
	visited := make(map[Node]bool)
	result := []Node{}

	// Recursive helper function
	var dfsVisit func(Node)
	dfsVisit = func(node Node) {
		visited[node] = true
		result = append(result, node)

		// Visit all adjacent nodes
		postNodes := node.GetPostNodes()
		for _, nextNode := range postNodes {
			if !visited[nextNode] {
				dfsVisit(nextNode)
			}
		}
	}

	dfsVisit(start)
	return result
}
