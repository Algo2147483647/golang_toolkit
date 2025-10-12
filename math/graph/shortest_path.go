package graph

// ShortestPath finds the shortest path between start and end nodes using BFS
// Returns the path as a slice of nodes, or an empty slice if no path exists
func ShortestPath(start, end Node) []Node {
	if start == end {
		return []Node{start}
	}

	visited := make(map[Node]bool)
	parent := make(map[Node]Node) // To reconstruct the path
	queue := []Node{start}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		// Check all adjacent nodes
		postNodes := current.GetPostNodes()
		for _, node := range postNodes {
			if !visited[node] {
				visited[node] = true
				parent[node] = current
				queue = append(queue, node)

				// If we reached the target node
				if node == end {
					// Reconstruct the path
					path := []Node{end}
					for n := end; parent[n] != start; n = parent[n] {
						path = append([]Node{parent[n]}, path...)
					}
					path = append([]Node{start}, path...)
					return path
				}
			}
		}
	}

	// No path found
	return []Node{}
}
