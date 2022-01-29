package thesaurus

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
)

type Thesaurus struct {
	Subjects Subjects `yaml:"subjects"`
}

type Subjects []*Subject

type Subject struct {
	ParentRelationship struct {
		PreferredParents    Terms `yaml:"preferredParents"`
		NonPreferredParents Terms `yaml:"nonPreferredParents"`
	} `yaml:"parentRelationship"`
	Term struct {
		PreferredTerms    Terms `yaml:"preferredTerms"`
		NonPreferredTerms Terms `yaml:"nonPreferredTerms"`
	} `yaml:"term"`
	DescriptiveNotes Notes `yaml:"descriptiveNotes"`
}

type Terms []*Term

func (t *Terms) IsEmpty() bool {
	return len(*t) < 1
}

func (t *Terms) First() *Term {
	for _, t := range *t {
		return t
	}
	return nil
}

type Term struct {
	TermText string `yaml:"termText"`
}

type Notes []*Note

type Note struct {
	NoteText string `yaml:"noteText"`
}

func NewThesaurus(filename string) (*Thesaurus, error) {
	t := &Thesaurus{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	if err := yaml.Unmarshal(b, t); err != nil {
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
