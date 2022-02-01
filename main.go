package main

import (
	"embed"
	"flag"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/thesaurus"
	"log"
	"os"
)

var (
	//go:embed assets
	assets  embed.FS
	builder *thesaurus.Builder
)

func init() {
	flag.Usage = usage
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
}

func usage() {
	if _, err := fmt.Fprintln(os.Stderr, "Usage: tb [flags]"); err != nil {
		log.Fatal(err)
	}
	flag.PrintDefaults()
}
