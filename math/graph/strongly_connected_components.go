package graph

// StronglyConnectedComponents finds all strongly connected components in a directed graph
// using Kosaraju's algorithm
func StronglyConnectedComponents(graph Graph) [][]Node {
	nodes := graph.GetNodes()

	// Step 1: Order the vertices in decreasing order of finish times in DFS
	visited := make(map[Node]bool)
	stack := []Node{}

	var dfsFill func(Node)
	dfsFill = func(node Node) {
		visited[node] = true
		postNodes := node.GetPostNodes()
		for _, neighbor := range postNodes {
			if !visited[neighbor] {
				dfsFill(neighbor)
			}
		}
		stack = append(stack, node)
	}

	// Fill vertices in stack according to their finish times
	for _, node := range nodes {
		if !visited[node] {
			dfsFill(node)
		}
	}

	// Step 2: Create a reversed graph
	// For this implementation, we'll assume that we can traverse the graph in reverse
	// by using GetPreNodes() instead of GetPostNodes()

	// Step 3: Process all vertices in order of decreasing finish times
	visited = make(map[Node]bool) // Reset visited
	components := [][]Node{}

	var dfsUtil func(Node, *[]Node)
	dfsUtil = func(node Node, component *[]Node) {
		visited[node] = true
		*component = append(*component, node)

		// Use GetPreNodes for the reversed graph
		preNodes := node.GetPreNodes()
		for _, neighbor := range preNodes {
			if !visited[neighbor] {
				dfsUtil(neighbor, component)
			}
		}
	}

	// Process vertices from stack
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if !visited[node] {
			component := []Node{}
			dfsUtil(node, &component)
			components = append(components, component)
		}
	}

	return components
}
