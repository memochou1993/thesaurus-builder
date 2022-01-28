package main

import (
	"github.com/memochou1993/thesaurus/thesaurus"
	"log"
)

var (
	t *thesaurus.Thesaurus
)

func main() {
	var err error
	if t, err = thesaurus.Parse("thesaurus.yaml"); err != nil {
		log.Fatal(err)
	}
	var root *thesaurus.Node
	if root, err = thesaurus.BuildTree(t.Subjects); err != nil {
		log.Fatal(err)
	}
	thesaurus.PrintTree(root, 0)
}
