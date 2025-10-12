package graph

type Node interface {
	GetPreNodes() []Node
	GetPostNodes() []Node
	GetEdges() []Edge
}
