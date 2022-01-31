package thesaurus

import (
	"errors"
	"fmt"
	"strings"
)

type Node struct {
	Subject  Subject
	Children []*Node
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}

func NewNode(subject Subject) *Node {
	return &Node{
		Subject: subject,
	}
}

func NewTree(t *Thesaurus) (root *Node, err error) {
	table := make(map[string]*Node, 1024)
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
			continue
		}
		table[preferredTerm.TermText] = nil
	}
	if root == nil {
		return nil, errors.New("root missing")
	}
	return root, buildTree(t.Subjects, table)
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

func PrintGraph(node *Node, level int) (s string) {
	preferredTerm := node.Subject.Term.PreferredTerms.First()
	s += fmt.Sprintf("%s|- %s\n", strings.Repeat("  ", level), preferredTerm.TermText)
	level++
	for _, child := range node.Children {
		s += PrintGraph(child, level)
	}
	return
}

func PrintJSON(node *Node) (s string) {
	s += "{"
	s += fmt.Sprintf("\"term\":{%s},", buildTermJSON(node))
	s += fmt.Sprintf("\"note\":{%s},", buildNoteJSON(node))
	s += "\"children\":["
	for i, child := range node.Children {
		s += PrintJSON(child)
		if i < len(node.Children)-1 {
			s += ","
		}
	}
	s += "]"
	s += "}"
	return
}

func buildTermJSON(node *Node) (s string) {
	s = "\"preferredTerms\":["
	preferredTerms := node.Subject.Term.PreferredTerms
	for i, term := range preferredTerms {
		s += fmt.Sprintf("{\"termText\":\"%s\"}", strings.ReplaceAll(term.TermText, "\"", "\\\""))
		if i < len(preferredTerms)-1 {
			s += ","
		}
	}
	s += "]"
	return
}

func buildNoteJSON(node *Node) (s string) {
	s = "\"descriptiveNotes\":["
	descriptiveNotes := node.Subject.Note.DescriptiveNotes
	for i, descriptiveNote := range descriptiveNotes {
		s += fmt.Sprintf("{\"noteText\":\"%s\"}", strings.ReplaceAll(descriptiveNote.NoteText, "\"", "\\\""))
		if i < len(descriptiveNotes)-1 {
			s += ","
		}
	}
	s += "]"
	return
}
