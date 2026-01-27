package rule_engine

// ExprNode is the base structure for expression nodes in the rule engine.
type ExprNode struct {
	Operator     OperatorIf             `json:"operator"`
	Params       map[string]interface{} `json:"params"`
	PreNodeList  []NodeIf               `json:"pre_node_list"`
	PostNodeList []NodeIf               `json:"post_node_list"`
}

func (n *ExprNode) GetType() string {
	return NodeTypeExpr
}

// Evaluate evaluates the expression node with given context (variables)
func (n *ExprNode) Evaluate() (result ValueIf, err error) {
	nodeListResult := make([]ValueIf, 0, len(n.PostNodeList))

	for _, node := range n.PostNodeList {
		nodeResult, err := Calculate(node)
		if err != nil {
			return nil, err
		}

		nodeListResult = append(nodeListResult, nodeResult)
	}

	return n.Operator.Evaluate(nodeListResult...)
}
