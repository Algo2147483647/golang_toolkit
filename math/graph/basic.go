package graph

func GetHeadTasks(g Graph) []Node {
	nodes := g.GetNodes()
	result := make([]Node, 0, len(nodes))

	for _, node := range nodes {
		if len(node.GetPreNodes()) == 0 {
			result = append(result, node)
		}
	}
	return result
}
