package thesaurus

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/helper"
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
	preferredTerm := node.Subject.Terms.FirstPreferred()
	s += fmt.Sprintf("\n%s|- %s", strings.Repeat("  ", level), preferredTerm.Text)
	level++
	for _, child := range node.Children {
		s += t.ToGraph(child, level)
	}
	return
}

func (t *Tree) toMD(node *Node, level int) (s string) {
	if level == 0 {
		s = fmt.Sprintf("%s\n===\n", t.Title)
		level++
	}
	preferredTerm := node.Subject.Terms.FirstPreferred()
	s += fmt.Sprintf("\n%s- %s\n", strings.Repeat("  ", level), preferredTerm.Text)
	for _, note := range node.Subject.Notes {
		s += fmt.Sprintf("\n%s  %s\n", strings.Repeat("  ", level), note.Text)
	}
	level++
	for _, child := range node.Children {
		s += t.toMD(child, level)
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

func NewTree(source *Resource) (thesaurus *Tree, err error) {
	helper.InitProgressBar(len(source.Subjects), "2/3", "Building thesaurus tree...")
	thesaurus = &Tree{
		Title: source.Title,
	}
	table := make(map[string]*Node, len(source.Subjects))
	for i, subject := range source.Subjects {
		preferredTerm := subject.Terms.FirstPreferred()
		if preferredTerm == nil {
			return nil, errors.New(fmt.Sprintf("preferred term missing (subject: #%d)", i+1))
		}
		if subject.ParentRelationships.FirstPreferred() == nil {
			if thesaurus.Root != nil {
				return nil, errors.New(fmt.Sprintf("preferred parent missing (subject: \"%s\")", preferredTerm.Text))
			}
			thesaurus.Root = NewNode(*subject)
			table[preferredTerm.Text] = thesaurus.Root
			if err := helper.ProgressBar.Add(1); err != nil {
				return nil, err
			}
			continue
		}
		table[preferredTerm.Text] = nil
	}
	if thesaurus.Root == nil {
		return nil, errors.New("root missing")
	}
	return thesaurus, buildTree(source.Subjects, table)
}

func buildTree(subjects Subjects, table map[string]*Node) (err error) {
	var orphans Subjects
	for i, subject := range subjects {
		if subject.ParentRelationships.FirstPreferred() == nil {
			continue
		}
		preferredTerm := subject.Terms.FirstPreferred()
		if preferredTerm == nil {
			return errors.New(fmt.Sprintf("preferred term missing (subject: #%d)", i+1))
		}
		preferredParent := subject.ParentRelationships.FirstPreferred()
		parent, ok := table[preferredParent.Text]
		if !ok {
			return errors.New(fmt.Sprintf("preferred parent missing (subject: \"%s\")", preferredTerm.Text))
		}
		if parent != nil {
			child := NewNode(*subject)
			parent.AppendChild(child)
			table[preferredTerm.Text] = child
			if err := helper.ProgressBar.Add(1); err != nil {
				return err
			}
			continue
		}
		orphans = append(orphans, subject)
	}
	if len(orphans) == len(subjects) {
		return
	}
	return buildTree(orphans, table)
}
