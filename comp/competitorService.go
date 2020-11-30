package comp

import (
	"errors"
	"fmt"
	"github.com/gregito/vrviewer/comp/model"
	"log"
	"strconv"

	"github.com/gregito/vrviewer/comp/dto"
)

func GetCompetitorOnCompetitionByCompetitionIdAndCompetitorName(competitionId int64, competitorName string) (*dto.CompetitorResult, error) {
	comp, err := GetCompetitionResultsByCompetitionId(competitionId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := findResultOfCompetitor(competitorName, *comp)
	sectionResult := findSectionResultOfCompetitor(competitorName, result)
	if sectionResult == nil {
		log.Println(fmt.Sprintf("Unable to find section result by the competitor's name (\"%s\")!", competitorName))
		return nil, errors.New("unable to find results for the given competitor")
	}
	var sections []dto.Section
	for _, s := range sectionResult {
		sections = append(sections, convertSectionResultAndSectionMapToSectionDto(s, comp.Sections))
	}
	return &dto.CompetitorResult{
		Name:            competitorName,
		CurrentPosition: strconv.FormatInt(result.Position, 10),
		SectionResults:  sections,
	}, nil
}

func findResultOfCompetitor(name string, cd model.CompetitionDetail) *model.Result {
	for _, p := range cd.Partitions {
		for _, r := range p.Results {
			if name == r.Name {
				return &r
			}
		}
	}
	return nil
}

func findSectionResultOfCompetitor(name string, cd *model.Result) []model.SectionResult {
	if cd == nil || cd.Name != name {
		return nil
	}
	return cd.SectionResults
}
