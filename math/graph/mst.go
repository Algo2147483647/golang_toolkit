package graph

import (
	"sort"
)

// EdgeWithWeight represents an edge with weight in a weighted graph
type EdgeWithWeight struct {
	From   Node
	To     Node
	Weight int
}

// WeightedGraph represents a weighted graph
type WeightedGraph interface {
	Graph
	GetWeightedEdges() []EdgeWithWeight
}

// DisjointSet represents a disjoint set data structure for Union-Find operations
type DisjointSet struct {
	parent map[Node]Node
	rank   map[Node]int
}

// NewDisjointSet creates a new disjoint set
func NewDisjointSet() *DisjointSet {
	return &DisjointSet{
		parent: make(map[Node]Node),
		rank:   make(map[Node]int),
	}
}

// MakeSet creates a new set containing only the given node
func (ds *DisjointSet) MakeSet(node Node) {
	ds.parent[node] = node
	ds.rank[node] = 0
}

// Find finds the representative (root) of the set containing the node
func (ds *DisjointSet) Find(node Node) Node {
	if ds.parent[node] != node {
		// Path compression
		ds.parent[node] = ds.Find(ds.parent[node])
	}
	return ds.parent[node]
}

// Union merges the sets containing node1 and node2
func (ds *DisjointSet) Union(node1, node2 Node) {
	root1 := ds.Find(node1)
	root2 := ds.Find(node2)

	if root1 != root2 {
		// Union by rank
		if ds.rank[root1] < ds.rank[root2] {
			ds.parent[root1] = root2
		} else if ds.rank[root1] > ds.rank[root2] {
			ds.parent[root2] = root1
		} else {
			ds.parent[root2] = root1
			ds.rank[root1]++
		}
	}
}

// KruskalMST finds the minimum spanning tree using Kruskal's algorithm
func KruskalMST(graph WeightedGraph) []EdgeWithWeight {
	edges := graph.GetWeightedEdges()
	nodes := graph.GetNodes()

	// Sort edges by weight
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].Weight < edges[j].Weight
	})

	// Initialize disjoint set
	ds := NewDisjointSet()
	for _, node := range nodes {
		ds.MakeSet(node)
	}

	// Initialize result
	mst := []EdgeWithWeight{}

	// Process edges in sorted order
	for _, edge := range edges {
		// If including this edge doesn't create a cycle
		if ds.Find(edge.From) != ds.Find(edge.To) {
			// Include this edge in MST
			mst = append(mst, edge)
			// Union the two sets
			ds.Union(edge.From, edge.To)
		}
	}

	return mst
}

// PrimMST finds the minimum spanning tree using Prim's algorithm
func PrimMST(graph WeightedGraph, start Node) []EdgeWithWeight {
	// This is a simplified version that assumes we can access all nodes and edges
	// A full implementation would require a priority queue

	// For now, let's implement a basic version
	edges := graph.GetWeightedEdges()
	nodes := graph.GetNodes()

	// If no start node specified, use the first node
	if start == nil && len(nodes) > 0 {
		start = nodes[0]
	}

	// Track visited nodes
	visited := make(map[Node]bool)
	visited[start] = true

	// Initialize MST
	mst := []EdgeWithWeight{}

	// Continue until all nodes are visited
	for len(mst) < len(nodes)-1 {
		minEdge := EdgeWithWeight{Weight: int(^uint(0) >> 1)} // Max int
		found := false

		// Find the minimum weight edge connecting visited and unvisited nodes
		for _, edge := range edges {
			// One node visited, one not visited
			if (visited[edge.From] && !visited[edge.To]) ||
				(visited[edge.To] && !visited[edge.From]) {
				if edge.Weight < minEdge.Weight {
					minEdge = edge
					found = true
				}
			}
		}

		// If no edge found, graph is not connected
		if !found {
			break
		}

		// Add the minimum edge to MST
		mst = append(mst, minEdge)

		// Mark the new node as visited
		if visited[minEdge.From] {
			visited[minEdge.To] = true
		} else {
			visited[minEdge.From] = true
		}
	}

	return mst
}

// MSTWeight calculates the total weight of the minimum spanning tree
func MSTWeight(mst []EdgeWithWeight) int {
	total := 0
	for _, edge := range mst {
		total += edge.Weight
	}
	return total
}
