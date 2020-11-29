package model

type Route struct {
	ID              int64       `json:"id"`
	DisplayPosition int64       `json:"displayPosition"`
	Label           string      `json:"label"`
	Color           Color       `json:"color"`
	Gender          RouteGender `json:"gender"`
	Options         string      `json:"options"`
}
