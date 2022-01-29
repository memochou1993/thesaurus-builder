package main

import (
	"fmt"
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
	if err = thesaurus.Build(thesaurus.PrintJSON(root), "dist"); err != nil {
		log.Fatal(err)
	}
	fmt.Println(thesaurus.PrintGraph(root, 0))
}
