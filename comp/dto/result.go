package dto

import "github.com/gregito/vrviewer/comp/model"

type CompetitorResult struct {
	CompetitionName     string
	CompetitionFinished bool
	Type                model.ClimbingType
	Name                string
	CurrentPosition     string
	SectionResults      []Section
}
