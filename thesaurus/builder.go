package thesaurus

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	TemplatePath = "template"
)

type Config struct {
	Data string
	Path string
}

func NewConfig(path string, data string) *Config {
	return &Config{
		Data: data,
		Path: path,
	}
}

func Build(config Config) error {
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

func copyHTML(config Config) error {
	filename := "index.html"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	if _, err = os.Stat(config.Path); os.IsNotExist(err) {
		if err = os.MkdirAll(config.Path, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.Path, filename)
	return ioutil.WriteFile(o, b, 0755)
}

func copyCSS(config Config) error {
	filename := "app.css"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	s := string(b)
	s = uglify(s, []string{"0 ", "px ", "title title-expandable"})
	if _, err = os.Stat(config.Path); os.IsNotExist(err) {
		if err = os.MkdirAll(config.Path, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.Path, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func copyJS(config Config) error {
	filename := "app.js"
	b, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", TemplatePath, filename))
	if err != nil {
		return err
	}
	s := string(b)
	s = uglify(s, []string{"const ", "let ", "title title-expandable", "title title-expanded"})
	s = strings.Replace(s, "\"__DATA__\"", config.Data, 1)
	if _, err = os.Stat(config.Path); os.IsNotExist(err) {
		if err = os.MkdirAll(config.Path, 0755); err != nil {
			return err
		}
	}
	o := fmt.Sprintf("%s/%s", config.Path, filename)
	return ioutil.WriteFile(o, []byte(s), 0755)
}

func uglify(s string, keywords []string) string {
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
