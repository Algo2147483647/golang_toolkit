package rule_engine

import "fmt"

type NodeIf interface {
	GetType() string
	Evaluate() (result ValueIf, err error)
}

const (
	NodeTypeExpr  = "expr_node"
	NodeTypeValue = "value_node"
)

// ExprNode is the base structure for expression nodes in the rule engine.
type NodeBase struct {
	Type         string
	Value        ValueIf                `json:"value"`
	Operator     OperatorIf             `json:"operator"`
	Params       map[string]interface{} `json:"params"`
	PreNodeList  []NodeIf               `json:"pre_node_list"`
	PostNodeList []NodeIf               `json:"post_node_list"`
}

func (n *NodeBase) GetType() string {
	return n.Type
}

// Evaluate evaluates the expression node with given context (variables)
func (n *NodeBase) Evaluate() (result ValueIf, err error) {
	switch n.GetType() {
	case NodeTypeExpr:
		return n.Value, nil

	case NodeTypeValue:
		nodeListResult := make([]ValueIf, 0, len(n.PostNodeList))

		for _, node := range n.PostNodeList {
			nodeResult, err := node.Evaluate()
			if err != nil {
				return nil, err
			}

			nodeListResult = append(nodeListResult, nodeResult)
		}

		return n.Operator.Evaluate(nodeListResult...)

	default:
		return nil, fmt.Errorf("unknown node type: %s", n.GetType())
	}
}
