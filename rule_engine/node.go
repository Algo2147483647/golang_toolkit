package rule_engine

import "fmt"

type NodeIf interface {
	GetType() string
}

const (
	NodeTypeExpr  = "expr_node"
	NodeTypeValue = "value_node"
)

func Calculate(node NodeIf) (result ValueIf, err error) {
	switch node.GetType() {
	case NodeTypeExpr:
		return node.(*ExprNode).Evaluate()

	case NodeTypeValue:
		return node.(*ValueNode).Value, nil
	}

	return nil, fmt.Errorf("unknown node type: %s", node.GetType())
}

type ValueNode struct {
	Value ValueIf `json:"value"`
}

func (n *ValueNode) GetType() string {
	return NodeTypeValue
}
