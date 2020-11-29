package comp

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gregito/vrviewer/webexec"
)

type Competition struct {
	ID   int64
	Name string
	Year int64
	Type ClimbingType
}

type InputElement struct {
	ID            int64          `json:"id"`
	Status        status         `json:"status"`
	ClimbingTypes []ClimbingType `json:"climbingTypes"`
	Year          int64          `json:"year"`
	Name          string         `json:"name"`
	Mnemonic      string         `json:"mnemonic"`
	Event         int64          `json:"event"`
	Notes         string         `json:"notes"`
}

type ClimbingType string

type status string

const (
	Active                   status       = "ACTIVE"
	Closed                   status       = "CLOSED"
	Boulder                  ClimbingType = "BOULDER"
	Lead                     ClimbingType = "LEAD"
	basePath                 string       = "https://vr2.mhssz.hu/api/1.0.0/competitions/"
	firstValidYear                        = 2018
	valueTodisableYearFilter              = 0
)

func GetCompetition(id int64) (*Competition, error) {
	result, err := webexec.ExecuteCall(fmt.Sprintf("%s%d", basePath, id), InputElement{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iep := convertInterfaceToInputElementPointer(result)
	cp := convertInputElementPointerToCompetitionPointer(iep)
	return cp, nil
}

func ListAllCompetitions() ([]*Competition, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, nil)
}

func ListCompetitionsByKind(kind *ClimbingType) ([]*Competition, error) {
	return ListCompetitionsByYearAndKind(valueTodisableYearFilter, kind)
}

func ListCompetitionsByYear(year int64) ([]*Competition, error) {
	return ListCompetitionsByYearAndKind(year, nil)
}

func ListCompetitionsByYearAndKind(year int64, kind *ClimbingType) ([]*Competition, error) {
	resp, err := webexec.ExecuteCall(basePath, []InputElement{})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	iaep := convertInterfaceArrayToInputElementPointerArray(resp)
	cp := convertInputElementArrayPointerToCompetitionPointerArray(iaep)
	result := filterCompetitions(cp, year, kind)
	return result, nil
}

func filterCompetitions(comps []*Competition, year int64, kind *ClimbingType) []*Competition {
	result := filterCompetitionsByYear(comps, year)
	if kind == nil || (*kind != Boulder && *kind != Lead) {
		return result
	}
	return filterCompetitionsByType(result, kind)
}

func filterCompetitionsByYear(comps []*Competition, year int64) []*Competition {
	if year != valueTodisableYearFilter {
		if year < firstValidYear {
			return []*Competition{}
		}
		var result []*Competition
		for _, c := range comps {
			if c.Year == year {
				result = append(result, c)
			}
		}
		return result
	} else {
		var result []*Competition
		for _, c := range comps {
			if c.Year >= firstValidYear {
				result = append(result, c)
			}
		}
		return result
	}
}

func filterCompetitionsByType(comps []*Competition, kind *ClimbingType) []*Competition {
	var result []*Competition
	for _, c := range comps {
		if c.Type == *kind {
			result = append(result, c)
		}
	}
	return result
}

func convertInterfaceArrayToInputElementPointerArray(intf interface{}) *[]InputElement {
	var i interface{} = intf
	c, ok := i.([]InputElement)
	if !ok {
		log.Printf("Unable to convert type (%s) to InputElement", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertInputElementArrayPointerToCompetitionPointerArray(fetchedData *[]InputElement) []*Competition {
	var result []*Competition
	if fetchedData != nil && len(*fetchedData) > 0 {
		for _, d := range *fetchedData {
			t := convertInputElementPointerToCompetitionPointer(&d)
			result = append(result, t)
		}
	}
	return result
}

func convertInterfaceToInputElementPointer(intf interface{}) *InputElement {
	var i interface{} = intf
	c, ok := i.(InputElement)
	if !ok {
		log.Printf("Unable to convert type (%s) to InputElement", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertInputElementPointerToCompetitionPointer(fetchedData *InputElement) *Competition {
	if fetchedData != nil && !reflect.ValueOf(&fetchedData).IsZero() {
		return &Competition{
			ID:   fetchedData.ID,
			Name: fetchedData.Name,
			Year: fetchedData.Year,
			Type: fetchedData.ClimbingTypes[0],
		}
	}
	return nil
}