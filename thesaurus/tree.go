package thesaurus

type Node struct {
	Subject  Subject
	Children []*Node
}

func NewNode(s Subject) *Node {
	return &Node{
		Subject: s,
	}
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}
