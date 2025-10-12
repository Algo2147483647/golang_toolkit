package graph

// HasCycle detects if there is a cycle in the graph using DFS
func HasCycle(graph Graph) bool {
	nodes := graph.GetNodes()
	visited := make(map[Node]bool)
	recursionStack := make(map[Node]bool)

	// Helper function for DFS
	var isCyclic func(Node) bool
	isCyclic = func(node Node) bool {
		// Mark the current node as visited and add to recursion stack
		visited[node] = true
		recursionStack[node] = true

		// Visit all neighbors
		postNodes := node.GetPostNodes()
		for _, neighbor := range postNodes {
			// If not visited, recursively visit it
			if !visited[neighbor] {
				if isCyclic(neighbor) {
					return true
				}
			} else if recursionStack[neighbor] {
				// If the node is in recursion stack, we found a cycle
				return true
			}
		}

		// Remove the node from recursion stack
		recursionStack[node] = false
		return false
	}

	// Check for cycles starting from each node
	for _, node := range nodes {
		if !visited[node] {
			if isCyclic(node) {
				return true
			}
		}
	}

	return false
}
