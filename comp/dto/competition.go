package dto

import "github.com/gregito/vrviewer/comp/model"

type Competition struct {
	ID   int64
	Name string             `header:"Name"`
	Year int64              `header:"Year"`
	Type model.ClimbingType `header:"Type"`
}

type BasicCompetition struct {
	Name   string             `header:"Name"`
	Type   model.ClimbingType `header:"Type"`
	Status string             `header:"Status"`
}
