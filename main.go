package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v2"
)

const (
	outputDir = "dist"
)

type Thesaurus struct {
	Subjects []Subject `yaml:"subjects"`
}

type Subject struct {
	ParentRelationship struct {
		PreferredParents    []Term `yaml:"preferredParents"`
		NonPreferredParents []Term `yaml:"nonPreferredParents"`
	} `yaml:"parentRelationship"`
	Term struct {
		PreferredTerms    []Term `yaml:"preferredTerms"`
		NonPreferredTerms []Term `yaml:"nonPreferredTerms"`
	} `yaml:"term"`
	DescriptiveNotes []Note `yaml:"descriptiveNotes"`
}

type Term struct {
	TermText string `yaml:"termText"`
}

type Note struct {
	NoteText string `yaml:"noteText"`
}

var (
	thesaurus Thesaurus
)

func main() {
	if err := parse(); err != nil {
		log.Fatal(err)
	}
	if err := print(); err != nil {
		log.Fatal(err)
	}
}

func parse() error {
	b, err := ioutil.ReadFile("thesaurus.yaml")
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(b, &thesaurus); err != nil {
		return err
	}
	for _, term := range thesaurus.Subjects {
		fmt.Println(term.Term.PreferredTerms)
	}
	return nil
}

func print() error {
	b, err := ioutil.ReadFile("template.html")
	if err != nil {
		return err
	}
	s := strings.Replace(string(b), "__BODY__", "hello", 1)
	if _, err = os.Stat(outputDir); os.IsNotExist(err) {
		if err = os.MkdirAll(outputDir, 0755); err != nil {
			return err
		}
	}
	filename := fmt.Sprintf("%s/index.html", outputDir)
	return ioutil.WriteFile(filename, []byte(s), 0755)
}
