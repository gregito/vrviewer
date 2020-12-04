package comp

import (
	"github.com/gregito/vrviewer/comp/common"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"strconv"
	"sync"
)

func GetCompetitorResults(name string, cds []model.CompetitionDetail) []dto.CompetitorResult {
	var result []dto.CompetitorResult
	wg := sync.WaitGroup{}
	for _, cd := range cds {
		wg.Add(1)
		go func(cd model.CompetitionDetail) {
			subResult := getCompetitorResultInCompetitionDetail(name, cd)
			if !common.IsStructEmpty(subResult) {
				result = append(result, subResult)
			}
			wg.Done()
		}(cd)
	}
	wg.Wait()
	return result
}

func getCompetitorResultInCompetitionDetail(name string, cd model.CompetitionDetail) dto.CompetitorResult {
	result := findResultOfCompetitor(name, cd)
	sectionResult := findSectionResultOfCompetitor(name, result)
	var sections []dto.Section
	climbingType := cd.Partitions[0].ClimbingType // hopefully, it won't break since MHSSZ doesn't organize competitions with multiple styles of climbing at the same time
	for _, s := range sectionResult {
		if isValidResult(climbingType, s) {
			sections = append(sections, convertSectionResultAndSectionMapToSectionDto(s, cd.Sections, climbingType))
		}
	}
	if len(sections) > 0 {
		return dto.CompetitorResult{
			CompetitionName: cd.Name,
			Type:            climbingType,
			Name:            name,
			CurrentPosition: strconv.FormatInt(result.Position, 10),
			SectionResults:  sections,
		}
	}
	return dto.CompetitorResult{}
}

func isValidResult(t model.ClimbingType, sr model.SectionResult) bool {
	if model.Lead == t {
		return sr.Points > 0
	}
	return !isInvalidBoulderResult(sr)
}

func isInvalidBoulderResult(sr model.SectionResult) bool {
	return sr.Tops == 0 && sr.TopTries == 0 && sr.Zones == 0 && sr.ZoneTries == 0
}

func findResultOfCompetitor(name string, cd model.CompetitionDetail) model.Result {
	for _, p := range cd.Partitions {
		for _, r := range p.Results {
			if name == r.Name && r.Position > 0 {
				return r
			}
		}
	}
	return model.Result{}
}

func findSectionResultOfCompetitor(name string, cd model.Result) []model.SectionResult {
	if cd.Name != name {
		var res []model.SectionResult
		return res
	}
	return cd.SectionResults
}
