package comp

import (
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
	"github.com/gregito/vrviewer/comp/model"
	"sync"
)

func GetCompetitorResults(name string, cds []model.CompetitionDetail) []dto.CompetitorResult {
	var result []dto.CompetitorResult
	crChan := make(chan *dto.CompetitorResult)
	wg := sync.WaitGroup{}
	for _, cd := range cds {
		wg.Add(1)
		go func(cd model.CompetitionDetail) {
			getCompetitorResultInCompetitionDetailIntoChannel(name, cd, crChan)
			wg.Done()
		}(cd)
	}
	go func() {
		wg.Wait()
		close(crChan)
	}()
	for r := range crChan {
		if r != nil {
			result = append(result, *r)
		}
	}
	return result
}

func getCompetitorResultInCompetitionDetailIntoChannel(name string, cd model.CompetitionDetail, crChan chan *dto.CompetitorResult) {
	crChan <- getCompetitorResultInCompetitionDetail(name, cd)
}

func getCompetitorResultInCompetitionDetail(name string, cd model.CompetitionDetail) *dto.CompetitorResult {
	result := findResultOfCompetitor(name, cd)
	sectionResult := findSectionResultOfCompetitor(name, result)
	var sections []dto.Section
	var climbingType *model.ClimbingType
	if len(cd.Partitions) == 0 {
		log.Println("Unable to determine competition climbing type because no Partition has been provided.\n")
	} else {
		climbingType = &cd.Partitions[0].ClimbingType
		for _, s := range sectionResult {
			if isValidResult(climbingType, s) {
				sections = append(sections, convertSectionResultAndSectionMapToSectionDto(s, cd.Sections, *climbingType))
			}
		}
		if len(sections) > 0 {
			return &dto.CompetitorResult{
				CompetitionName:     cd.Name,
				Type:                *climbingType,
				Name:                name,
				CompetitionFinished: cd.IsFinished(),
				CurrentPosition:     result.Position,
				SectionResults:      sections,
			}
		}
	}
	return nil
}

func isValidResult(t *model.ClimbingType, sr model.SectionResult) bool {
	if t == nil {
		log.Println("Section result is not valid since no climbing type has given to decide its validity upon")
		return false
	}
	if model.Lead == *t {
		return sr.HasValidLeadResult()
	}
	return sr.HasValidBoulderResult()
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
