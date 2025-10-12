package graph

type Edge interface {
	GetNodes() (Node, Node)
}
