package thesaurus

import (
	"errors"
	"fmt"
)

type Node struct {
	Subject  Subject
	Children []*Node
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}

func NewNode(s Subject) *Node {
	return &Node{
		Subject: s,
	}
}

func BuildTree(subjects []Subject) (root *Node, err error) {
	table := make(map[string]*Node, 1024)
	for _, subject := range subjects {
		if len(subject.ParentRelationship.PreferredParents) == 0 {
			if len(subject.Term.PreferredTerms) == 0 {
				return nil, errors.New("invalid root node")
			}
			root = NewNode(subject)
			table[subject.Term.PreferredTerms[0].TermText] = root
			break
		}
	}
	if root == nil {
		return nil, errors.New("invalid root node")
	}
	knit(subjects, table)
	return root, nil
}

func knit(subjects []Subject, table map[string]*Node) {
	var remaining []Subject
	for _, subject := range subjects {
		if len(subject.ParentRelationship.PreferredParents) == 0 {
			continue
		}
		if len(subject.Term.PreferredTerms) == 0 {
			continue
		}
		if node, ok := table[subject.ParentRelationship.PreferredParents[0].TermText]; ok {
			child := NewNode(subject)
			node.AppendChild(child)
			table[subject.Term.PreferredTerms[0].TermText] = child
			continue
		}
		remaining = append(remaining, subject)
	}
	if len(remaining) == len(subjects) || len(remaining) == 0 {
		return
	}
	knit(remaining, table)
}

func PrintTree(node *Node, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	fmt.Printf("%s|- %s\n", indent, node.Subject.Term.PreferredTerms[0].TermText)
	level++
	for _, node := range node.Children {
		PrintTree(node, level)
	}
}
