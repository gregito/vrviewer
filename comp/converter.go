package comp

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
)

func convertInterfaceArrayToCompetitionPointerArray(intf interface{}) *[]model.Competition {
	var i interface{} = intf
	c, ok := i.([]model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertCompetitionArrayPointerToCompetitionPointerArray(fetchedData *[]model.Competition) []*dto.Competition {
	var result []*dto.Competition
	if fetchedData != nil && len(*fetchedData) > 0 {
		for _, d := range *fetchedData {
			t := convertCompetitionPointerToCompetitionPointer(&d)
			result = append(result, t)
		}
	}
	return result
}

func convertInterfaceToCompetitionPointer(intf interface{}) *model.Competition {
	var i interface{} = intf
	c, ok := i.(model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertInterfaceToCompetitionDetailPointer(intf interface{}) *model.CompetitionDetail {
	var i interface{} = intf
	c, ok := i.(model.CompetitionDetail)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.CompetitionDetail", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertCompetitionPointerToCompetitionPointer(fetchedData *model.Competition) *dto.Competition {
	if fetchedData != nil && !reflect.ValueOf(&fetchedData).IsZero() {
		return &dto.Competition{
			ID:   fetchedData.ID,
			Name: fetchedData.Name,
			Year: fetchedData.Year,
			Type: fetchedData.ClimbingTypes[0],
		}
	}
	return nil
}
