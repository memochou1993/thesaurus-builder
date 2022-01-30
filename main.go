package main

import (
	"fmt"
	"github.com/memochou1993/thesaurus/thesaurus"
	"log"
)

var (
	config *thesaurus.Config
)

func init() {
	thesaurus.ParseFlags(config)
}

func main() {
	var t *thesaurus.Thesaurus
	var err error
	if t, err = thesaurus.NewThesaurus(config.File); err != nil {
		log.Fatal(err)
	}
	var root *thesaurus.Node
	if root, err = thesaurus.NewTree(t.Subjects); err != nil {
		log.Fatal(err)
	}
	config.Data = thesaurus.PrintJSON(root)
	if err = thesaurus.Build(config); err != nil {
		log.Fatal(err)
	}
	fmt.Println(thesaurus.PrintGraph(root, 0))
}
