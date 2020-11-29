package comp

import (
	"fmt"
	"log"
	"reflect"

	"github.com/gregito/vrviewer/comp/model"
)

func convertInterfaceArrayToInputElementPointerArray(intf interface{}) *[]model.Competition {
	var i interface{} = intf
	c, ok := i.([]model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertInputElementArrayPointerToCompetitionPointerArray(fetchedData *[]model.Competition) []*CompetitionDto {
	var result []*CompetitionDto
	if fetchedData != nil && len(*fetchedData) > 0 {
		for _, d := range *fetchedData {
			t := convertInputElementPointerToCompetitionPointer(&d)
			result = append(result, t)
		}
	}
	return result
}

func convertInterfaceToInputElementPointer(intf interface{}) *model.Competition {
	var i interface{} = intf
	c, ok := i.(model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", intf))
		return nil
	}
	return &c
}

func convertInputElementPointerToCompetitionPointer(fetchedData *model.Competition) *CompetitionDto {
	if fetchedData != nil && !reflect.ValueOf(&fetchedData).IsZero() {
		return &CompetitionDto{
			ID:   fetchedData.ID,
			Name: fetchedData.Name,
			Year: fetchedData.Year,
			Type: fetchedData.ClimbingTypes[0],
		}
	}
	return nil
}
