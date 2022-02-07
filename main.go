package main

import (
	"embed"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/thesaurus"
	"log"
)

var (
	//go:embed themes
	themesDir embed.FS
	builder   *thesaurus.Builder
)

func init() {
	builder = thesaurus.NewBuilder()
	builder.SetDefaultThemesDir(themesDir)
	builder.Init()
}

func main() {
	var r *thesaurus.Resource
	var t *thesaurus.Tree
	var err error
	if r, err = thesaurus.NewResource(builder.Filename); err != nil {
		log.Fatal(err)
	}
	if t, err = thesaurus.NewTree(r); err != nil {
		log.Fatal(err)
	}
	if err = builder.Build(t); err != nil {
		log.Fatal(err)
	}
	fmt.Println(t.ToGraph(t.Root, 0))
}
