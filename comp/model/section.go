package model

type Section struct {
	ID       int64   `json:"id"`
	Rank     int64   `json:"rank"`
	Name     string  `json:"name"`
	TryLimit int64   `json:"tryLimit"`
	Routes   []Route `json:"routes"`
}
