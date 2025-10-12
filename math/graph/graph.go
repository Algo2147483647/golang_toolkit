package graph

type Graph interface {
	GetNodes() []Node
	GetEdges() []Edge
}
