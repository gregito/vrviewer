package model

type CompetitionDetail struct {
	ID         int64              `json:"id"`
	Name       string             `json:"name"`
	Status     string             `json:"status"`
	Sections   map[string]Section `json:"sections"`
	Partitions []Partition        `json:"partitions"`
}
