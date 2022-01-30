package thesaurus

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	TemplatePath = "assets"
)

type Config struct {
	Data      string
	File      string
	OutputDir string
}

func ParseFlags(config *Config) {
	flag.StringVar(&config.File, "f", "thesaurus.yaml", "thesaurus file")
	flag.StringVar(&config.OutputDir, "o", "dist", "output directory")
	flag.Parse()
}

func Build(config *Config) error {
	if err := copyHTML(config); err != nil {
		return err
	}
	if err := copyCSS(config); err != nil {
		return err
	}
	if err := copyJS(config); err != nil {
		return err
	}
	return nil
}

func copyHTML(config *Config) error {
	filename := "index.html"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	if _, err = os.Stat(config.OutputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(config.OutputDir, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.OutputDir, filename)
	return ioutil.WriteFile(o, b, 0755)
}

func copyCSS(config *Config) error {
	filename := "app.css"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	s := string(b)
	s = minify(s, []string{"0 ", "px ", "title title-expandable"})
	if _, err = os.Stat(config.OutputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(config.OutputDir, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.OutputDir, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func copyJS(config *Config) error {
	filename := "app.js"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	s := string(b)
	s = minify(s, []string{"const ", "let ", "title title-expandable", "title title-expanded"})
	s = strings.Replace(s, "\"__DATA__\"", config.Data, 1)
	if _, err = os.Stat(config.OutputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(config.OutputDir, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.OutputDir, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func minify(s string, keywords []string) string {
	for _, k := range keywords {
		s = strings.ReplaceAll(s, k, strings.ReplaceAll(k, " ", "_"))
	}
	s = strings.ReplaceAll(s, " ", "")
	for _, k := range keywords {
		s = strings.ReplaceAll(s, strings.ReplaceAll(k, " ", "_"), k)
	}
	s = strings.ReplaceAll(s, "\n", "")
	return s
}
