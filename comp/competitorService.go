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
			subResult := GetCompetitorResultInCompetitionDetail(name, cd)
			if !common.IsStructEmpty(subResult) {
				result = append(result, subResult)
			}
			wg.Done()
		}(cd)
	}
	wg.Wait()
	return result
}

func GetCompetitorResultInCompetitionDetail(name string, cd model.CompetitionDetail) dto.CompetitorResult {
	result := findResultOfCompetitor(name, cd)
	sectionResult := findSectionResultOfCompetitor(name, result)
	var sections []dto.Section
	for _, s := range sectionResult {
		sections = append(sections, convertSectionResultAndSectionMapToSectionDto(s, cd.Sections))
	}
	return dto.CompetitorResult{
		CompetitionName: cd.Name,
		Name:            name,
		CurrentPosition: strconv.FormatInt(result.Position, 10),
		SectionResults:  sections,
	}
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
