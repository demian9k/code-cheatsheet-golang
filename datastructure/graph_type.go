package datastructure

type NodePair struct {
	FromNode int
	ToNode   int
}

type GraphRel struct {
	NodePair
	Weight int
}

func NewNodePair(fromNode int, toNode int) *NodePair {
	return &NodePair{ToNode: toNode, FromNode: fromNode}
}

func NewGraphRel(fromNode int, toNode int, weight int) *GraphRel {
	return &GraphRel{NodePair: *NewNodePair(toNode, fromNode), Weight: weight}
}
