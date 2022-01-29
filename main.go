package main

import (
	"bytes"
	"encoding/json"
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
	printGraph(root)
	printJSON(root)
}

func printGraph(root *thesaurus.Node) {
	fmt.Println(thesaurus.PrintGraph(root, 0))
}

func printJSON(root *thesaurus.Node) {
	var dst bytes.Buffer
	if err := json.Indent(&dst, []byte(thesaurus.PrintJSON(root)), "", "  "); err != nil {
		log.Fatal(err)
	}
	fmt.Println(dst.String())
}
