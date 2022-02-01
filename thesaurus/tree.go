package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/helper"
	"github.com/schollz/progressbar/v3"
	"log"
	"strings"
)

type Tree struct {
	Title string `json:"title"`
	Root  *Node  `json:"root"`
}

func (t *Tree) ToJSON() string {
	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}
	return string(b)
}

func (t *Tree) ToGraph(node *Node, level int) (s string) {
	if level == 0 {
		s = t.Title
		level++
	}
	preferredTerm := node.Subject.Term.PreferredTerms.First()
	s += fmt.Sprintf("\n%s|- %s", strings.Repeat("  ", level), preferredTerm.TermText)
	level++
	for _, child := range node.Children {
		s += t.ToGraph(child, level)
	}
	return
}

type Node struct {
	Subject  Subject `json:"subject"`
	Children []*Node `json:"children,omitempty"`
}

func (n *Node) AppendChild(node *Node) {
	n.Children = append(n.Children, node)
}

func NewNode(subject Subject) *Node {
	return &Node{
		Subject: subject,
	}
}

func NewTree(source *Source) (thesaurus *Tree, err error) {
	bar := helper.NewProgressBar(len(source.Subjects), "2/3", "Building thesaurus tree...")
	thesaurus = &Tree{
		Title: source.Title,
	}
	table := make(map[string]*Node, len(source.Subjects))
	for i, subject := range source.Subjects {
		if subject.Term.PreferredTerms.IsEmpty() {
			return nil, errors.New(fmt.Sprintf("preferred term missing (subject: #%d)", i+1))
		}
		preferredTerm := subject.Term.PreferredTerms.First()
		if subject.ParentRelationship.PreferredParents.IsEmpty() {
			if thesaurus.Root != nil {
				return nil, errors.New(fmt.Sprintf("preferred parent missing (subject: \"%s\")", preferredTerm.TermText))
			}
			thesaurus.Root = NewNode(*subject)
			table[preferredTerm.TermText] = thesaurus.Root
			if err := bar.Add(1); err != nil {
				return nil, err
			}
			continue
		}
		table[preferredTerm.TermText] = nil
	}
	if thesaurus.Root == nil {
		return nil, errors.New("root missing")
	}
	return thesaurus, buildTree(source.Subjects, table, bar)
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
