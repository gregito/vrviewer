package model

type RouteResult struct {
	Route          int64   `json:"route"`
	BestHoldNumber int64   `json:"bestHoldNumber"`
	Tries          []int64 `json:"tries"`
}
