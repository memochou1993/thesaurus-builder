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
	builder.InitFlags()
}

func main() {
	var s *thesaurus.Source
	var t *thesaurus.Tree
	var err error
	if s, err = thesaurus.NewSource(builder.Filename); err != nil {
		log.Fatal(err)
	}
	if t, err = thesaurus.NewTree(s); err != nil {
		log.Fatal(err)
	}
	if err = builder.Build(t); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.ToGraph(t.Root, 0))
}
