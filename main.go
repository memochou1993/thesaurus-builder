package main

import (
	"github.com/memochou1993/thesaurus/thesaurus"
	"log"
)

var (
	source *thesaurus.Thesaurus
	root   *thesaurus.Node
)

func main() {
	var err error
	source, err = thesaurus.Parse("source.yaml")
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range source.Subjects {
		subject := s
		if len(subject.ParentRelationship.PreferredParents) == 0 {
			root = thesaurus.NewNode(subject)
			continue
		}
		if len(subject.Term.PreferredTerms) == 0 {
			continue
		}
		if subject.ParentRelationship.PreferredParents[0].TermText == root.Subject.Term.PreferredTerms[0].TermText {
			root.AppendChild(thesaurus.NewNode(subject))
			continue
		}
	}

	for _, term := range root.Children {
		log.Println("term:", term.Subject.Term.PreferredTerms[0].TermText)
	}
}
