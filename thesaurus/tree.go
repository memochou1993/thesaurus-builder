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
	if root, err = buildRoot(subjects, table); err != nil {
		return nil, err
	}
	if err = buildBranch(subjects, table); err != nil {
		return nil, err
	}
	return root, nil
}

func buildRoot(subjects []Subject, table map[string]*Node) (root *Node, err error) {
	for _, subject := range subjects {
		if len(subject.ParentRelationship.PreferredParents) == 0 {
			if len(subject.Term.PreferredTerms) == 0 {
				return nil, errors.New("invalid root")
			}
			root = NewNode(subject)
			table[subject.Term.PreferredTerms[0].TermText] = root
			return
		}
	}
	return nil, errors.New("root missing")
}

func buildBranch(subjects []Subject, table map[string]*Node) (err error) {
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
	if len(remaining) == len(subjects) {
		if len(remaining) != 0 {
			for _, subject := range remaining {
				for _, parent := range subject.ParentRelationship.PreferredParents {
					return errors.New(fmt.Sprintf("parent missing: %s", parent.TermText))
				}
			}
		}
		return
	}
	return buildBranch(remaining, table)
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
