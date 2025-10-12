package graph

// TreeDiameter calculates the diameter of a tree (the longest path between any two nodes)
// The tree is assumed to be connected and undirected
func TreeDiameter(graph Graph) int {
	nodes := graph.GetNodes()
	if len(nodes) == 0 {
		return 0
	}

	// First BFS from any node to find one end of the diameter
	firstEnd := nodes[0]
	_, farthestNode := bfsFarthest(firstEnd)

	// Second BFS from the farthest node to find the other end of the diameter
	distance, _ := bfsFarthest(farthestNode)

	return distance
}

// bfsFarthest performs BFS and returns the maximum distance and the farthest node
func bfsFarthest(start Node) (int, Node) {
	visited := make(map[Node]bool)
	queue := []nodeDistancePair{{node: start, distance: 0}}
	visited[start] = true

	maxDistance := 0
	farthestNode := start

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.distance > maxDistance {
			maxDistance = current.distance
			farthestNode = current.node
		}

		// Add all unvisited neighbors
		preNodes := current.node.GetPreNodes()
		postNodes := current.node.GetPostNodes()

		// Combine pre and post nodes (for undirected tree)
		allNeighbors := append(preNodes, postNodes...)

		for _, neighbor := range allNeighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, nodeDistancePair{node: neighbor, distance: current.distance + 1})
			}
		}
	}

	return maxDistance, farthestNode
}

// nodeDistancePair is a helper struct to store a node and its distance from the start node
type nodeDistancePair struct {
	node     Node
	distance int
}
