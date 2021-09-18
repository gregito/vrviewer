package comp

import (
	"strings"
	"sync"

	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
	"github.com/gregito/vrviewer/comp/model"
)

func GetCompetitorResults(name string, cds []model.CompetitionDetail) []dto.CompetitorResult {
	var result []dto.CompetitorResult
	crChan := make(chan *dto.CompetitorResult)
	wg := sync.WaitGroup{}
	log.Printf("About to fetch competition results for: %s\n", name)
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
	crChan <- getCompetitorResultInCompetitionDetailV2(name, cd)
}

func getCompetitorResultInCompetitionDetailV2(name string, cd model.CompetitionDetail) *dto.CompetitorResult {
	if len(cd.Partitions) == 0 {
		log.Println("Unable to determine competition climbing type because no Partition has been provided.\n")
		return nil
	} else {
		var climbingType *model.ClimbingType
		climbingType = &cd.Partitions[0].ClimbingType
		var agr []dto.AgeGroupResult
		competitorResults := findResultsOfCompetitor(name, cd)
		for _, result := range competitorResults {
			var sections []dto.Section
			var sectionResults []model.SectionResult
			for _, sectionResult := range findSectionResultOfCompetitor(name, result) {
				sectionResults = append(sectionResults, sectionResult)
			}
			for _, s := range sectionResults {
				if isValidResult(climbingType, s) {
					sections = append(sections, convertSectionResultAndSectionMapToSectionDto(s, cd.Sections, *climbingType))
				}
			}
			if len(sections) > 0 {
				agr = append(agr, dto.AgeGroupResult{
					AgeGroup:      result.AgeGroup,
					FinalPosition: result.Result.Position,
					Results:       sections,
				})
			}
		}
		if len(competitorResults) > 0 && len(agr) > 0 {
			return &dto.CompetitorResult{
				CompetitionName:     cd.Name,
				Category:            *climbingType,
				CompetitionFinished: cd.IsFinished(),
				AgeGroupResult:      agr,
			}
		}
		return nil
	}
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

func findResultsOfCompetitor(name string, cd model.CompetitionDetail) []dto.ExtendedResult {
	results := make([]dto.ExtendedResult, 0)
	for _, p := range cd.Partitions {
		for _, r := range p.Results {
			if name == strings.TrimSpace(r.Name) && r.Position > 0 {
				results = append(results, dto.ExtendedResult{
					AgeGroup: p.AgeGroup,
					Result:   *trimWhitespacesFromNameInResult(&r),
				})
			}
		}
	}
	return results
}

func trimWhitespacesFromNameInResult(r *model.Result) *model.Result {
	r.Name = strings.TrimSpace(r.Name)
	return r
}

func findSectionResultOfCompetitor(name string, cd dto.ExtendedResult) []model.SectionResult {
	if cd.Result.Name != name {
		var res []model.SectionResult
		return res
	}
	return cd.Result.SectionResults
}
