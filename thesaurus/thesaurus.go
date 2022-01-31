package thesaurus

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Thesaurus struct {
	Subjects Subjects `yaml:"subjects"`
}

type Subjects []*Subject

type Subject struct {
	Term struct {
		PreferredTerms    Terms `yaml:"preferredTerms"`
		NonPreferredTerms Terms `yaml:"nonPreferredTerms"`
	} `yaml:"term"`
	ParentRelationship struct {
		PreferredParents    Terms `yaml:"preferredParents"`
		NonPreferredParents Terms `yaml:"nonPreferredParents"`
	} `yaml:"parentRelationship"`
	Note struct {
		DescriptiveNotes Notes `yaml:"descriptiveNotes"`
	} `yaml:"note"`
}

type Terms []*Term

func (t *Terms) IsEmpty() bool {
	return len(*t) < 1
}

func (t *Terms) First() *Term {
	for _, term := range *t {
		return term
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

func NewThesaurus(filename string) (t *Thesaurus, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	if err = yaml.Unmarshal(b, &t); err != nil {
		return
	}
	return t, nil
}
