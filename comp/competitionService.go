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
	valueToDisableYearFilter        = 0
)

func GetCompetition(id int64) (dto.Competition, error) {
	result, err := webexec.ExecuteCall(fmt.Sprintf("%s%d", basePath, id), model.Competition{})
	if err != nil {
		log.Println(err)
		return dto.Competition{}, err
	}
	iep := convertInterfaceToCompetitionPointer(result)
	cp := convertCompetitionPointerToCompetitionPointer(iep)
	return cp, nil
}

func ListAllCompetitionsSimplified() []dto.Competition {
	var ct model.ClimbingType
	return ListCompetitionsByYearAndKind(valueToDisableYearFilter, ct)
}

func ListCompetitionsByKind(kind model.ClimbingType) []dto.Competition {
	return ListCompetitionsByYearAndKind(valueToDisableYearFilter, kind)
}

func ListCompetitionsByYear(year int64) []dto.Competition {
	var ct model.ClimbingType
	return ListCompetitionsByYearAndKind(year, ct)
}

func GetCompetitionResultsByCompetitionId(id int64) (model.CompetitionDetail, error) {
	resp, err := webexec.ExecuteCall(fmt.Sprintf("%s%d/results", basePath, id), model.CompetitionDetail{})
	if err != nil {
		log.Println(err)
		return model.CompetitionDetail{}, err
	}
	result := convertInterfaceToCompetitionDetailPointer(resp)
	return result, nil
}

func ListCompetitionsByYearAndKind(year int64, kind model.ClimbingType) []dto.Competition {
	resp, err := webexec.ExecuteCall(basePath, []model.Competition{})
	if err != nil {
		log.Println(err)
		var empty []dto.Competition
		return empty
	}
	iaep := convertInterfaceArrayToCompetitionArray(resp)
	cp := convertCompetitionArrayToCompetitionArray(iaep)
	result := filterCompetitions(cp, year, kind)
	return result
}

func filterCompetitions(comps []dto.Competition, year int64, kind model.ClimbingType) []dto.Competition {
	result := collectCompetitionsByYear(comps, year)
	if kind != model.Boulder && kind != model.Lead {
		return result
	}
	return collectCompetitionsByType(result, kind)
}

func collectCompetitionsByYear(comps []dto.Competition, year int64) []dto.Competition {
	var result []dto.Competition
	if year != valueToDisableYearFilter {
		if year < firstValidYear {
			return result
		}
		for _, c := range comps {
			if c.Year == year {
				result = append(result, c)
			}
		}
		return result
	} else {
		for _, c := range comps {
			if c.Year >= firstValidYear {
				result = append(result, c)
			}
		}
		return result
	}
}

func collectCompetitionsByType(comps []dto.Competition, kind model.ClimbingType) []dto.Competition {
	var result []dto.Competition
	for _, c := range comps {
		if c.Type == kind {
			result = append(result, c)
		}
	}
	return result
}
