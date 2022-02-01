package thesaurus

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	AssetsPath = "assets"
)

type Builder struct {
	Assets    embed.FS
	Filename  string
	OutputDir string
	Root      *Node
}

func (b *Builder) SetAssets(assets embed.FS) {
	b.Assets = assets
}

func (b *Builder) SetRoot(root *Node) {
	b.Root = root
}

func (b *Builder) InitFlags() {
	flag.Usage = func() {
		_, _ = fmt.Fprintln(os.Stderr, "Usage: tb [flags]")
		flag.PrintDefaults()
	}
	flag.StringVar(&b.Filename, "f", "thesaurus.yaml", "source file")
	flag.StringVar(&b.OutputDir, "o", "dist", "output directory")
	flag.Parse()
}

func (b *Builder) Build(root *Node) error {
	bar := NewProgressBar(10000, "3/3", "Generating thesaurus assets...")
	go StartPermanentProgress(bar)
	defer FinishPermanentProgress(bar)
	b.SetRoot(root)
	if err := b.makeDir(); err != nil {
		return err
	}
	if err := b.copyHTML(); err != nil {
		return err
	}
	if err := b.copyCSS(); err != nil {
		return err
	}
	if err := b.copyJS(); err != nil {
		return err
	}
	if err := b.copyJSON(); err != nil {
		return err
	}
	return nil
}

func (b *Builder) makeDir() error {
	if _, err := os.Stat(b.OutputDir); os.IsNotExist(err) {
		return os.MkdirAll(b.OutputDir, 0755)
	}
	return nil
}

func (b *Builder) copyHTML() error {
	filename := "index.html"
	data, err := b.Assets.ReadFile(fmt.Sprintf("%s/%s", AssetsPath, filename))
	if err != nil {
		return err
	}
	o := fmt.Sprintf("%s/%s", b.OutputDir, filename)
	return ioutil.WriteFile(o, data, 0755)
}

func (b *Builder) copyCSS() error {
	filename := "app.css"
	data, err := b.Assets.ReadFile(fmt.Sprintf("%s/%s", AssetsPath, filename))
	if err != nil {
		return err
	}
	protectedKeywords := []string{"0 ", "px ", "title title-expandable"}
	s := string(data)
	s = minify(s, protectedKeywords)
	o := fmt.Sprintf("%s/%s", b.OutputDir, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func (b *Builder) copyJS() error {
	filename := "app.js"
	data, err := b.Assets.ReadFile(fmt.Sprintf("%s/%s", AssetsPath, filename))
	if err != nil {
		return err
	}
	protectedKeywords := []string{"async ", "await ", "const ", "let ", "title title-expandable", "title title-expanded"}
	s := string(data)
	s = minify(s, protectedKeywords)
	o := fmt.Sprintf("%s/%s", b.OutputDir, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func (b *Builder) copyJSON() error {
	filename := "data.json"
	o := fmt.Sprintf("%s/%s", b.OutputDir, filename)
	return ioutil.WriteFile(o, []byte(b.Root.ToJSON()), 0755)
}

func NewBuilder() *Builder {
	return &Builder{}
}

func minify(s string, protectedKeywords []string) string {
	for _, k := range protectedKeywords {
		s = strings.ReplaceAll(s, k, strings.ReplaceAll(k, " ", "_"))
	}
	s = strings.ReplaceAll(s, " ", "")
	for _, k := range protectedKeywords {
		s = strings.ReplaceAll(s, strings.ReplaceAll(k, " ", "_"), k)
	}
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
