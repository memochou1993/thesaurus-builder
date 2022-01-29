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

func NewNode(subject Subject) *Node {
	return &Node{
		Subject: subject,
	}
}

func NewTree(subjects Subjects) (root *Node, err error) {
	table := make(map[string]*Node, 1024)
	for i, subject := range subjects {
		if subject.Term.PreferredTerms.IsEmpty() {
			return nil, errors.New(fmt.Sprintf("preferred term missing (subject: #%d)", i+1))
		}
		preferredTerm := subject.Term.PreferredTerms.First()
		if subject.ParentRelationship.PreferredParents.IsEmpty() {
			if root != nil {
				return nil, errors.New(fmt.Sprintf("preferred parent missing (subject: \"%s\")", preferredTerm.TermText))
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
	return root, buildTree(subjects, table)
}

func buildTree(subjects Subjects, table map[string]*Node) (err error) {
	var orphans Subjects
	for i, subject := range subjects {
		if subject.ParentRelationship.PreferredParents.IsEmpty() {
			continue
		}
		if subject.Term.PreferredTerms.IsEmpty() {
			return errors.New(fmt.Sprintf("preferred term missing (subject: #%d)", i+1))
		}
		preferredParent := subject.ParentRelationship.PreferredParents.First()
		parent, ok := table[preferredParent.TermText]
		if !ok {
			preferredTerm := subject.Term.PreferredTerms.First()
			return errors.New(fmt.Sprintf("preferred parent missing (subject: \"%s\")", preferredTerm.TermText))
		}
		if parent != nil {
			child := NewNode(*subject)
			parent.AppendChild(child)
			preferredTerm := subject.Term.PreferredTerms.First()
			table[preferredTerm.TermText] = child
			continue
		}
		orphans = append(orphans, subject)
	}
	if len(orphans) == len(subjects) {
		return
	}
	return buildTree(orphans, table)
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
