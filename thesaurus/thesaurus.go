package thesaurus

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
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

func NewThesaurus() *Thesaurus {
	return &Thesaurus{}
}

func Parse(filename string) (*Thesaurus, error) {
	t := NewThesaurus()
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, &t); err != nil {
		return nil, err
	}
	return t, nil
}

func Build(path string) error {
	b, err := ioutil.ReadFile("template.html")
	if err != nil {
		return err
	}
	s := strings.Replace(string(b), "__BODY__", "hello", 1)
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	filename := fmt.Sprintf("%s/index.html", path)
	return ioutil.WriteFile(filename, []byte(s), 0755)
}
