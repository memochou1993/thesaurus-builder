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

func (n *Node) Contains(node *Node) bool {
	for _, child := range n.Children {
		for _, childTerm := range child.Subject.Term.PreferredTerms {
			for _, nodeTerm := range node.Subject.Term.PreferredTerms {
				if childTerm.TermText == nodeTerm.TermText {
					return true
				}
			}
		}
	}
	return false
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
			for _, term := range subject.Term.PreferredTerms {
				table[term.TermText] = root
				break
			}
			return
		}
	}
	return nil, errors.New("root missing")
}

func buildBranch(subjects []Subject, table map[string]*Node) (err error) {
	var orphans []Subject
	for _, subject := range subjects {
		if len(subject.ParentRelationship.PreferredParents) == 0 {
			continue
		}
		if len(subject.Term.PreferredTerms) == 0 {
			continue
		}
		isOrphan := false
		for _, parent := range subject.ParentRelationship.PreferredParents {
			if parent, ok := table[parent.TermText]; ok {
				child := NewNode(subject)
				if !parent.Contains(child) {
					parent.AppendChild(child)
				}
				for _, term := range subject.Term.PreferredTerms {
					if _, ok := table[term.TermText]; !ok {
						table[term.TermText] = child
					}
					break
				}
				continue
			}
			isOrphan = true
		}
		if isOrphan {
			orphans = append(orphans, subject)
		}
	}
	if len(orphans) == len(subjects) {
		if len(orphans) != 0 {
			for _, subject := range orphans {
				for _, parent := range subject.ParentRelationship.PreferredParents {
					if _, ok := table[parent.TermText]; !ok {
						return errors.New(fmt.Sprintf("parent missing: %s", parent.TermText))
					}
				}
			}
		}
		return
	}
	return buildBranch(orphans, table)
}

func PrintTree(node *Node, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	for _, term := range node.Subject.Term.PreferredTerms {
		fmt.Printf("%s|- %s\n", indent, term.TermText)
		break
	}
	level++
	for _, child := range node.Children {
		PrintTree(child, level)
	}
}
