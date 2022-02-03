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
	Terms               Terms `json:"terms" yaml:"terms"`
	ParentRelationships Terms `json:"parentRelationships,omitempty" yaml:"parentRelationships"`
	Notes               Notes `json:"notes,omitempty" yaml:"notes"`
}

type Terms []*Term

func (t *Terms) FirstPreferred() *Term {
	for _, term := range *t {
		if term.Preferred {
			return term
		}
	}
	return nil
}

type Term struct {
	Text      string `json:"text" yaml:"text"`
	Preferred bool   `json:"preferred" yaml:"preferred"`
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
