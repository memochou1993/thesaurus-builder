package thesaurus

import (
	"errors"
	"fmt"
)

type Node struct {
	Subject  Subject
	Children []*Node
}

func (n *Node) FirstPreferredParent() *Term {
	return n.Subject.ParentRelationship.PreferredParents.First()
}

func (n *Node) FirstPreferredTerm() *Term {
	return n.Subject.Term.PreferredTerms.First()
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) ContainsChild(node *Node) bool {
	for _, child := range n.Children {
		if child.FirstPreferredParent().TermText == node.FirstPreferredTerm().TermText {
			return true
		}
	}
	return false
}

func NewNode(subject Subject) *Node {
	return &Node{
		Subject: subject,
	}
}

func BuildTree(subjects Subjects) (root *Node, err error) {
	table := make(map[string]*Node, 1024)
	if root, err = buildRoot(subjects, table); err != nil {
		return nil, err
	}
	if err = buildBranch(subjects, table); err != nil {
		return nil, err
	}
	return root, nil
}

func buildRoot(subjects Subjects, table map[string]*Node) (root *Node, err error) {
	for i, subject := range subjects {
		if subject.Term.PreferredTerms.IsEmpty() {
			return nil, errors.New(fmt.Sprintf("preferred term missing (subject #%d)", i+1))
		}
		preferredTerm := subject.Term.PreferredTerms.First()
		if subject.ParentRelationship.PreferredParents.IsEmpty() {
			if root != nil {
				term := subject.Term.PreferredTerms.First()
				return nil, errors.New(fmt.Sprintf("preferred parent missing (subject \"%s\")", term.TermText))
			}
			root = NewNode(*subject)
			table[preferredTerm.TermText] = root
			continue
		}
		table[preferredTerm.TermText] = nil
	}
	if root == nil {
		return nil, errors.New("root missing")
	}
	return
}

func buildBranch(subjects Subjects, table map[string]*Node) (err error) {
	var orphans Subjects
	for i, subject := range subjects {
		if subject.ParentRelationship.PreferredParents.IsEmpty() {
			continue
		}
		if subject.Term.PreferredTerms.IsEmpty() {
			return errors.New(fmt.Sprintf("preferred term missing (subject #%d)", i+1))
		}
		preferredParent := subject.ParentRelationship.PreferredParents.First()
		parent, ok := table[preferredParent.TermText]
		if !ok {
			term := subject.Term.PreferredTerms.First()
			return errors.New(fmt.Sprintf("preferred parent missing (subject \"%s\")", term.TermText))
		}
		if parent != nil {
			child := NewNode(*subject)
			if !parent.ContainsChild(child) {
				parent.AppendChild(child)
			}
			preferredTerm := subject.Term.PreferredTerms.First()
			table[preferredTerm.TermText] = child
			continue
		}
		orphans = append(orphans, subject)
	}
	if len(orphans) == len(subjects) {
		return
	}
	return buildBranch(orphans, table)
}

func PrintTree(node *Node, level int) {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	preferredTerm := node.Subject.Term.PreferredTerms.First()
	fmt.Printf("%s|- %s\n", indent, preferredTerm.TermText)
	level++
	for _, child := range node.Children {
		PrintTree(child, level)
	}
}
