package dto

import "github.com/gregito/vrviewer/comp/model"

type CompetitorResult struct {
	CompetitionName string
	Type            model.ClimbingType
	Name            string
	CurrentPosition string
	SectionResults  []Section
}
