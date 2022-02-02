package thesaurus

import (
	"github.com/memochou1993/thesaurus-builder/helper"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Resource struct {
	Title    string   `json:"title" yaml:"title"`
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
	if len(*t) == 0 {
		return true
	}
	for _, term := range *t {
		if term.Text == "" {
			return true
		}
	}
	return false
}

func (t *Terms) First() *Term {
	for _, term := range *t {
		return term
	}
	return nil
}

type Term struct {
	Text string `json:"text" yaml:"text"`
}

type Notes []*Note

type Note struct {
	Text string `json:"text" yaml:"text"`
}

func NewResource(filename string) (r *Resource, err error) {
	var b []byte
	b, err = ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	go helper.StartPermanentProgress(1200, "1/3", "Unmarshalling thesaurus file...")
	defer helper.FinishPermanentProgress()
	if err = yaml.Unmarshal(b, &r); err != nil {
		return
	}
	return r, nil
}
