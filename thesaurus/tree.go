package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/schollz/progressbar/v3"
	"log"
	"strings"
)

type Node struct {
	Subject  Subject `json:"subject"`
	Children []*Node `json:"children,omitempty"`
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}

func (n *Node) ToJSON() string {
	b, err := json.Marshal(n)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (n *Node) ToGraph(level int) (s string) {
	if level == 0 {
		s = "\nThesaurus Tree"
		level++
	}
	preferredTerm := n.Subject.Term.PreferredTerms.First()
	s += fmt.Sprintf("\n%s|- %s", strings.Repeat("  ", level), preferredTerm.TermText)
	level++
	for _, child := range n.Children {
		s += child.ToGraph(level)
	}
	return
}

func NewNode(subject Subject) *Node {
	return &Node{
		Subject: subject,
	}
}

func NewTree(t *Thesaurus) (root *Node, err error) {
	bar := NewProgressBar(len(t.Subjects), "2/3", "Building thesaurus tree...")
	table := make(map[string]*Node, len(t.Subjects))
	for i, subject := range t.Subjects {
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
			if err := bar.Add(1); err != nil {
				return nil, err
			}
			continue
		}
		table[preferredTerm.TermText] = nil
	}
	if root == nil {
		return nil, errors.New("root missing")
	}
	return root, buildTree(t.Subjects, table, bar)
}

func buildTree(subjects Subjects, table map[string]*Node, bar *progressbar.ProgressBar) (err error) {
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
			if err := bar.Add(1); err != nil {
				return err
			}
			continue
		}
		orphans = append(orphans, subject)
	}
	if len(orphans) == len(subjects) {
		return
	}
	return buildTree(orphans, table, bar)
}
