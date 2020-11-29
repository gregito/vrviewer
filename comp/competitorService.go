package comp

import (
	"log"

	"github.com/gregito/vrviewer/comp/dto"
)

func GetCompetitorResultOnCompetition(competitionId int64, competitorName string) (*dto.CompetitorResult, error) {
	_, err := GetCompetitionResultsByCompetitionId(competitionId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &dto.CompetitorResult{}, nil
}
