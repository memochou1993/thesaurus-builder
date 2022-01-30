package main

import (
	"embed"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/thesaurus"
	"log"
)

var (
	//go:embed assets
	assets  embed.FS
	builder *thesaurus.Builder
)

func init() {
	builder = thesaurus.NewBuilder()
	builder.SetAssets(assets)
	builder.ParseFlags()
}

func main() {
	var t *thesaurus.Thesaurus
	var root *thesaurus.Node
	var err error
	if t, err = thesaurus.NewThesaurus(builder.Filename); err != nil {
		log.Fatal(err)
	}
	if root, err = thesaurus.NewTree(t); err != nil {
		log.Fatal(err)
	}
	if err = builder.Build(root); err != nil {
		log.Fatal(err)
	}
	fmt.Println(thesaurus.PrintGraph(root, 0))
}
