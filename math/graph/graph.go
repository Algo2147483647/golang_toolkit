package graph

type Graph interface {
	GetNodes() []Node
	GetEdges() []Edge
}

type Node interface {
	GetPreNodes() []Node
	GetPostNodes() []Node
}

type Edge interface {
	GetNodes() (Node, Node)
}
