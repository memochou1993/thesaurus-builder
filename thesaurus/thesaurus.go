package thesaurus

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Thesaurus struct {
	Subjects Subjects `json:"subjects" yaml:"subjects"`
}

type Subjects []*Subject

type Subject struct {
	Term struct {
		PreferredTerms    Terms `json:"preferredTerms" yaml:"preferredTerms"`
		NonPreferredTerms Terms `json:"nonPreferredTerms,omitempty" yaml:"nonPreferredTerms"`
	} `json:"term" yaml:"term"`
	ParentRelationship struct {
		PreferredParents    Terms `json:"preferredParents" yaml:"preferredParents"`
		NonPreferredParents Terms `json:"nonPreferredParents,omitempty" yaml:"nonPreferredParents"`
	} `json:"parentRelationship" yaml:"parentRelationship"`
	Note struct {
		DescriptiveNotes Notes `json:"descriptiveNotes,omitempty" yaml:"descriptiveNotes"`
	} `json:"note" yaml:"note"`
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
	TermText string `json:"termText" yaml:"termText"`
}

type Notes []*Note

type Note struct {
	NoteText string `json:"noteText" yaml:"noteText"`
}

func NewThesaurus(filename string) (t *Thesaurus, err error) {
	bar := NewProgressBar(100000, "1/3", "Unmarshalling thesaurus file...")
	go StartPermanentProgress(bar)
	defer FinishPermanentProgress(bar)
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
