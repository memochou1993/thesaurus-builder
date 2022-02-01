package main

import (
	"embed"
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
	builder.InitFlags()
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
}
