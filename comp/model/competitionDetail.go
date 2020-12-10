package model

type CompetitionDetail struct {
	ID         int64              `json:"id"`
	Name       string             `json:"name"`
	Status     string             `json:"status"`
	Sections   map[string]Section `json:"sections"`
	Partitions []Partition        `json:"partitions"`
}

func (cd CompetitionDetail) IsEmpty() bool {
	return cd.ID == 0 && len(cd.Name) == 0 && len(cd.Status) == 0 && cd.Sections == nil && cd.Partitions == nil
}

func (cd CompetitionDetail) IsFinished() bool {
	return cd.Status == "CLOSED"
}
