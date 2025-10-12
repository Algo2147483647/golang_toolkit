package graph

// LowestCommonAncestor finds the lowest common ancestor of two nodes in a tree
// Assumes that the tree is rooted and nodes have parent pointers or a common root is known
func LowestCommonAncestor(root, node1, node2 Node) Node {
	if root == nil || node1 == nil || node2 == nil {
		return nil
	}

	// If either node is the root, then root is the LCA
	if node1 == root || node2 == root {
		return root
	}

	// Find paths from root to both nodes
	path1 := pathToNode(root, node1)
	path2 := pathToNode(root, node2)

	if path1 == nil || path2 == nil {
		return nil // One or both nodes not found
	}

	// Find the last common node in both paths
	var lca Node
	i := 0
	for i < len(path1) && i < len(path2) && path1[i] == path2[i] {
		lca = path1[i]
		i++
	}

	return lca
}

// pathToNode finds the path from root to target node using DFS
func pathToNode(root, target Node) []Node {
	if root == nil {
		return nil
	}

	path := []Node{}
	if findPathDFS(root, target, &path) {
		return path
	}

	return nil // Target not found
}

// findPathDFS is a helper function that performs DFS to find the path to target
func findPathDFS(current, target Node, path *[]Node) bool {
	if current == nil {
		return false
	}

	// Add current node to path
	*path = append(*path, current)

	// If target found
	if current == target {
		return true
	}

	// Check children
	postNodes := current.GetPostNodes()
	for _, child := range postNodes {
		if findPathDFS(child, target, path) {
			return true
		}
	}

	// If target not found in this subtree, remove current node from path
	*path = (*path)[:len(*path)-1]
	return false
}

// LCAWithParentPointers finds LCA when nodes have parent pointers
// This assumes that each node has a way to access its parent, which would need to be
// implemented in the concrete Node type
func LCAWithParentPointers(node1, node2 Node) Node {
	if node1 == nil || node2 == nil {
		return nil
	}

	// Create sets to store ancestors
	ancestors1 := make(map[Node]bool)

	// Store all ancestors of node1
	current := node1
	for current != nil {
		ancestors1[current] = true
		// Here we would traverse to parent, but since we don't have parent access in the interface,
		// we'll assume there's some way to get the parent in a real implementation
		current = nil // Placeholder
	}

	// Traverse ancestors of node2 to find the first one that is also an ancestor of node1
	current = node2
	for current != nil {
		if ancestors1[current] {
			return current
		}
		// Here we would traverse to parent
		current = nil // Placeholder
	}

	return nil
}
