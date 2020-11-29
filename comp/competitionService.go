package comp

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"github.com/gregito/vrviewer/webexec"
)

const (
	basePath                 string = "https://vr2.mhssz.hu/api/1.0.0/competitions/"
	firstValidYear                  = 2018
	valueTodisableYearFilter        = 0
)

func GetCompetition(id int64) (*dto.Competition, error) {
	result, err := webexec.ExecuteCall(fmt.Sprintf("%s%d", basePath, id), model.Competition{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iep := convertInterfaceToCompetitionPointer(result)
	cp := convertCompetitionPointerToCompetitionPointer(iep)
	return cp, nil
}

func ListAllCompetitions() ([]*dto.Competition, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, nil)
}

func ListCompetitionsByKind(kind *model.ClimbingType) ([]*dto.Competition, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, kind)
}

func ListCompetitionsByYear(year int64) ([]*dto.Competition, error) {
	return ListCompetitionsByYearAndKind(year, nil)
}

func GetCompetitionResultsByCompetitionId(id int64) (*model.CompetitionDetail, error) {
	resp, err := webexec.ExecuteCall(basePath, model.CompetitionDetail{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result := convertInterfaceToCompetitionDetailPointer(resp)
	return result, nil
}

func ListCompetitionsByYearAndKind(year int64, kind *model.ClimbingType) ([]*dto.Competition, error) {
	resp, err := webexec.ExecuteCall(basePath, []model.Competition{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iaep := convertInterfaceArrayToCompetitionPointerArray(resp)
	cp := convertCompetitionArrayPointerToCompetitionPointerArray(iaep)
	result := filterCompetitions(cp, year, kind)
	return result, nil
}

func filterCompetitions(comps []*dto.Competition, year int64, kind *model.ClimbingType) []*dto.Competition {
	result := filterCompetitionsByYear(comps, year)
	if kind == nil || (*kind != model.Boulder && *kind != model.Lead) {
		return result
	}
	return filterCompetitionsByType(result, kind)
}

func filterCompetitionsByYear(comps []*dto.Competition, year int64) []*dto.Competition {
	if year != valueTodisableYearFilter {
		if year < firstValidYear {
			return []*dto.Competition{}
		}
		var result []*dto.Competition
		for _, c := range comps {
			if c.Year == year {
				result = append(result, c)
			}
		}
		return result
	} else {
		var result []*dto.Competition
		for _, c := range comps {
			if c.Year >= firstValidYear {
				result = append(result, c)
			}
		}
		return result
	}
}

func filterCompetitionsByType(comps []*dto.Competition, kind *model.ClimbingType) []*dto.Competition {
	var result []*dto.Competition
	for _, c := range comps {
		if c.Type == *kind {
			result = append(result, c)
		}
	}
	return result
}
