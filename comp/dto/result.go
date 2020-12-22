package dto

import "github.com/gregito/vrviewer/comp/model"

type CompetitorResult struct {
	CompetitionName     string
	Year                int64
	Category            model.ClimbingType
	CompetitionFinished bool
	AgeGroupResult      []AgeGroupResult
}
