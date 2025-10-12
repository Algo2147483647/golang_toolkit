package graph

// SubtreeSize calculates the size of the subtree rooted at each node in a tree
// Returns a map where key is the node and value is the size of its subtree
func SubtreeSize(root Node) map[Node]int {
	sizes := make(map[Node]int)

	if root == nil {
		return sizes
	}

	// Calculate subtree sizes using post-order traversal
	calculateSubtreeSize(root, sizes)

	return sizes
}

// calculateSubtreeSize is a helper function that recursively calculates subtree sizes
func calculateSubtreeSize(node Node, sizes map[Node]int) int {
	if node == nil {
		return 0
	}

	size := 1 // Count the node itself

	// Recursively calculate sizes of children subtrees
	children := node.GetPostNodes()
	for _, child := range children {
		size += calculateSubtreeSize(child, sizes)
	}

	// Store the size for this node
	sizes[node] = size

	return size
}

// SubtreeSizeWithLeafCount calculates both subtree sizes and leaf counts
func SubtreeSizeWithLeafCount(root Node) (map[Node]int, map[Node]int) {
	sizes := make(map[Node]int)
	leafCounts := make(map[Node]int)

	if root == nil {
		return sizes, leafCounts
	}

	calculateSubtreeAndLeafCount(root, sizes, leafCounts)

	return sizes, leafCounts
}

// calculateSubtreeAndLeafCount is a helper that calculates both subtree sizes and leaf counts
func calculateSubtreeAndLeafCount(node Node, sizes map[Node]int, leafCounts map[Node]int) (int, int) {
	if node == nil {
		return 0, 0
	}

	// If this is a leaf node
	children := node.GetPostNodes()
	if len(children) == 0 {
		sizes[node] = 1
		leafCounts[node] = 1
		return 1, 1
	}

	size := 1   // Count the node itself
	leaves := 0 // Count of leaf nodes in subtree

	// Recursively calculate for children
	for _, child := range children {
		childSize, childLeaves := calculateSubtreeAndLeafCount(child, sizes, leafCounts)
		size += childSize
		leaves += childLeaves
	}

	sizes[node] = size
	leafCounts[node] = leaves

	return size, leaves
}
