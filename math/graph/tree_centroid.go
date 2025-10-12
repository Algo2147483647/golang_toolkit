package graph

// TreeCentroid finds the centroid(s) of a tree
// A centroid is a node whose removal splits the tree into subtrees each having at most n/2 nodes
func TreeCentroid(root Node) []Node {
	if root == nil {
		return []Node{}
	}

	// First, calculate subtree sizes
	subtreeSizes := SubtreeSize(root)
	totalNodes := subtreeSizes[root]

	centroids := []Node{}
	findCentroids(root, nil, subtreeSizes, totalNodes, &centroids)

	return centroids
}

// findCentroids is a helper function that recursively finds centroids
func findCentroids(node, parent Node, subtreeSizes map[Node]int, totalNodes int, centroids *[]Node) {
	isCentroid := true

	// Check if the subtree rooted at this node has more than half of all nodes
	// (only if this is not the root)
	if parent != nil && subtreeSizes[node] > totalNodes/2 {
		isCentroid = false
	}

	// Check all children
	children := node.GetPostNodes()
	for _, child := range children {
		// Check if the subtree of child has more than n/2 nodes when we remove current node
		if totalNodes-subtreeSizes[child] > totalNodes/2 {
			isCentroid = false
		}

		// Continue DFS
		findCentroids(child, node, subtreeSizes, totalNodes, centroids)
	}

	// Check the "parent" component (the rest of the tree excluding this node's subtree)
	if parent != nil && totalNodes-subtreeSizes[node] > totalNodes/2 {
		isCentroid = false
	}

	if isCentroid {
		*centroids = append(*centroids, node)
	}
}

// TreeCentroidAlternative finds centroids using an alternative approach
// by removing leaves layer by layer until 1 or 2 nodes remain
func TreeCentroidAlternative(graph Graph) []Node {
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

	// Repeatedly remove leaf nodes (degree = 1) until we have 1 or 2 nodes left
	for len(remainingNodes) > 2 {
		// Find all leaf nodes
		leaves := []Node{}
		for node := range remainingNodes {
			if degree[node] == 1 { // Leaf node has degree 1
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

	// The remaining nodes are the centroids
	centroids := []Node{}
	for node := range remainingNodes {
		centroids = append(centroids, node)
	}

	return centroids
}
