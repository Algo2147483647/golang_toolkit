package graph

// TreeCenter finds the center node(s) of a tree
// The center of a tree is the node (or two nodes) that minimizes the maximum distance to all other nodes
func TreeCenter(graph Graph) []Node {
	nodes := graph.GetNodes()
	if len(nodes) == 0 {
		return []Node{}
	}

	// Make a copy of all nodes which we'll modify
	remainingNodes := make(map[Node]bool)
	for _, node := range nodes {
		remainingNodes[node] = true
	}

	// Count the degree of each node
	degree := make(map[Node]int)
	for node := range remainingNodes {
		preNodes := node.GetPreNodes()
		postNodes := node.GetPostNodes()
		degree[node] = len(preNodes) + len(postNodes)
	}

	// Repeatedly remove leaf nodes until we have 1 or 2 nodes left
	for len(remainingNodes) > 2 {
		// Find all leaf nodes (degree <= 1)
		leaves := []Node{}
		for node := range remainingNodes {
			if degree[node] <= 1 {
				leaves = append(leaves, node)
			}
		}

		// Remove all leaf nodes
		for _, leaf := range leaves {
			delete(remainingNodes, leaf)

			// Update degree of neighbors
			preNodes := leaf.GetPreNodes()
			postNodes := leaf.GetPostNodes()
			neighbors := append(preNodes, postNodes...)

			for _, neighbor := range neighbors {
				if remainingNodes[neighbor] {
					degree[neighbor]--
				}
			}
		}
	}

	// The remaining nodes are the center(s)
	centers := []Node{}
	for node := range remainingNodes {
		centers = append(centers, node)
	}

	return centers
}
