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
	config := thesaurus.NewConfig("dist", thesaurus.PrintJSON(root))
	if err = thesaurus.Build(*config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(thesaurus.PrintGraph(root, 0))
}
