package main

type Thesaurus struct {
	Subjects []Subject `json:"subjects"`
}

type Subject struct {
	PreferredParents    interface{} `json:"preferredParents"`
	NonPreferredParents interface{} `json:"nonPreferredParents"`
	PreferredTerms      []struct {
		TermText string `json:"termText"`
	} `json:"preferredTerms"`
	NonPreferredTerms []struct {
		TermText string `json:"termText"`
	} `json:"nonPreferredTerms"`
	DescriptiveNotes []struct {
		NoteText string `json:"noteText"`
	} `json:"descriptiveNotes"`
}
