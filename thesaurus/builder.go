package thesaurus

import (
	"embed"
	"flag"
	"fmt"
	"github.com/memochou1993/thesaurus-builder/helper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	DefaultAssetsPath = "assets"
	DefaultAssetHTML  = "index.html"
	DefaultAssetCSS   = "style.css"
	DefaultAssetJS    = "main.js"
	DefaultAssetJSON  = "data.json"
)

type Builder struct {
	AssetsDir        string
	DefaultAssetsDir embed.FS
	Filename         string
	OutputDir        string
	Tree             *Tree
}

func (b *Builder) SetDefaultAssetsDir(d embed.FS) {
	b.DefaultAssetsDir = d
}

func (b *Builder) SetTree(t *Tree) {
	b.Tree = t
}

func (b *Builder) Init() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: tb [flags]")
		flag.PrintDefaults()
	}
	flag.StringVar(&b.AssetsDir, "a", "", "assets directory")
	flag.StringVar(&b.Filename, "f", "thesaurus.yaml", "thesaurus file")
	flag.StringVar(&b.OutputDir, "o", "dist", "output directory")
	flag.Parse()
	if b.AssetsDir != "" {
		b.checkAssetsDir()
	}
}

func (b *Builder) Build(t *Tree) (err error) {
	go helper.StartPermanentProgress(1200, "3/3", "Generating thesaurus assets...")
	defer helper.FinishPermanentProgress()
	b.SetTree(t)
	if err = b.makeOutputDir(); err != nil {
		return
	}
	if err = b.writeHTML(); err != nil {
		return
	}
	if err = b.writeCSS(); err != nil {
		return
	}
	if err = b.writeJS(); err != nil {
		return
	}
	if err = b.writeJSON(); err != nil {
		return
	}
	return
}

func (b *Builder) checkAssetsDir() {
	if _, err := os.Stat(b.AssetsDir); os.IsNotExist(err) {
		log.Fatal(err)
	}
}

func (b *Builder) makeOutputDir() error {
	if _, err := os.Stat(b.OutputDir); err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(b.OutputDir, 0755)
		}
		return err
	}
	return nil
}

func (b *Builder) writeHTML() error {
	data, err := b.readAsset(DefaultAssetHTML)
	if err != nil {
		return err
	}
	s := string(data)
	s = strings.Replace(s, "__TITLE__", b.Tree.Title, 1)
	return b.writeAsset(DefaultAssetHTML, []byte(s))
}

func (b *Builder) writeCSS() error {
	data, err := b.readAsset(DefaultAssetCSS)
	if err != nil {
		return err
	}
	keywords := []string{"0 ", "px "}
	return b.writeAsset(DefaultAssetCSS, minify(data, keywords))
}

func (b *Builder) writeJS() error {
	data, err := b.readAsset(DefaultAssetJS)
	if err != nil {
		return err
	}
	keywords := []string{"async ", "await ", "const ", "let ", "term term-expandable", "term term-expanded"}
	return b.writeAsset(DefaultAssetJS, minify(data, keywords))
}

func (b *Builder) writeJSON() error {
	return b.writeAsset(DefaultAssetJSON, []byte(b.Tree.ToJSON()))
}

func (b *Builder) readAsset(filename string) ([]byte, error) {
	if b.AssetsDir != "" {
		b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", b.AssetsDir, filename))
		if err == nil {
			return b, err
		}
		if !os.IsNotExist(err) {
			log.Fatal(err)
		}
	}
	return b.DefaultAssetsDir.ReadFile(fmt.Sprintf("%s/%s", DefaultAssetsPath, filename))
}

func (b *Builder) writeAsset(filename string, data []byte) error {
	return ioutil.WriteFile(fmt.Sprintf("%s/%s", b.OutputDir, filename), data, 0755)
}

func NewBuilder() *Builder {
	return &Builder{}
}

func minify(b []byte, keywords []string) []byte {
	s := string(b)
	for _, k := range keywords {
		s = strings.ReplaceAll(s, k, strings.ReplaceAll(k, " ", "_"))
	}
	s = strings.ReplaceAll(s, " ", "")
	for _, k := range keywords {
		s = strings.ReplaceAll(s, strings.ReplaceAll(k, " ", "_"), k)
	}
	s = strings.ReplaceAll(s, "\n", "")
	return []byte(s)
}
