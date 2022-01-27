package main

type Thesaurus struct {
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	ParentRelationship struct {
		PreferredParents    []Term `json:"preferredParents"`
		NonPreferredParents []Term `json:"nonPreferredParents"`
	} `json:"parentRelationship"`
	Term struct {
		PreferredTerms    []Term `json:"preferredTerms"`
		NonPreferredTerms []Term `json:"nonPreferredTerms"`
	} `json:"term"`
	DescriptiveNotes []Note `json:"descriptiveNotes"`
}

type Term struct {
	TermText string `json:"termText"`
}

type Note struct {
	NoteText string `json:"noteText"`
}
