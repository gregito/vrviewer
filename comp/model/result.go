package model

type Result struct {
	Participation  int64           `json:"participation"`
	StartNumber    int64           `json:"startNumber"`
	Name           string          `json:"name"`
	Position       int64           `json:"position"`
	Tags           string          `json:"tags"`
	SectionResults []SectionResult `json:"sectionResults"`
}
