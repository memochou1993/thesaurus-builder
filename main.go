package main

import (
	"github.com/memochou1993/thesaurus/thesaurus"
	"log"
)

func main() {
	var t *thesaurus.Thesaurus
	var err error
	if t, err = thesaurus.NewThesaurus("thesaurus.yaml"); err != nil {
		log.Fatal(err)
	}
	var root *thesaurus.Node
	if root, err = thesaurus.NewTree(t.Subjects); err != nil {
		log.Fatal(err)
	}
	thesaurus.PrintTree(root, 0)
}
