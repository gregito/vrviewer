package comp

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp/model"
	"github.com/gregito/vrviewer/webexec"
)

type CompetitionDto struct {
	ID   int64
	Name string
	Year int64
	Type model.ClimbingType
}

const (
	basePath                 string = "https://vr2.mhssz.hu/api/1.0.0/competitions/"
	firstValidYear                  = 2018
	valueTodisableYearFilter        = 0
)

func GetCompetition(id int64) (*CompetitionDto, error) {
	result, err := webexec.ExecuteCall(fmt.Sprintf("%s%d", basePath, id), model.Competition{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iep := convertInterfaceToInputElementPointer(result)
	cp := convertInputElementPointerToCompetitionPointer(iep)
	return cp, nil
}

func ListAllCompetitions() ([]*CompetitionDto, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, nil)
}

func ListCompetitionsByKind(kind *model.ClimbingType) ([]*CompetitionDto, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, kind)
}

func ListCompetitionsByYear(year int64) ([]*CompetitionDto, error) {
	return ListCompetitionsByYearAndKind(year, nil)
}

func ListCompetitionsByYearAndKind(year int64, kind *model.ClimbingType) ([]*CompetitionDto, error) {
	resp, err := webexec.ExecuteCall(basePath, []model.Competition{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iaep := convertInterfaceArrayToInputElementPointerArray(resp)
	cp := convertInputElementArrayPointerToCompetitionPointerArray(iaep)
	result := filterCompetitions(cp, year, kind)
	return result, nil
}

func filterCompetitions(comps []*CompetitionDto, year int64, kind *model.ClimbingType) []*CompetitionDto {
	result := filterCompetitionsByYear(comps, year)
	if kind == nil || (*kind != model.Boulder && *kind != model.Lead) {
		return result
	}
	return filterCompetitionsByType(result, kind)
}

func filterCompetitionsByYear(comps []*CompetitionDto, year int64) []*CompetitionDto {
	if year != valueTodisableYearFilter {
		if year < firstValidYear {
			return []*CompetitionDto{}
		}
		var result []*CompetitionDto
		for _, c := range comps {
			if c.Year == year {
				result = append(result, c)
			}
		}
		return result
	} else {
		var result []*CompetitionDto
		for _, c := range comps {
			if c.Year >= firstValidYear {
				result = append(result, c)
			}
		}
		return result
	}
}

func filterCompetitionsByType(comps []*CompetitionDto, kind *model.ClimbingType) []*CompetitionDto {
	var result []*CompetitionDto
	for _, c := range comps {
		if c.Type == *kind {
			result = append(result, c)
		}
	}
	return result
}
