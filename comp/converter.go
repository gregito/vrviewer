package comp

import (
	"fmt"
	"github.com/gregito/vrviewer/comp/common"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"log"
)

func convertInterfaceArrayToCompetitionArray(source interface{}) []model.Competition {
	var i = source
	c, ok := i.([]model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", source))
		return nil
	}
	return c
}

func convertCompetitionArrayToCompetitionArray(fetchedData []model.Competition) []dto.Competition {
	var result []dto.Competition
	if fetchedData != nil && len(fetchedData) > 0 {
		for _, d := range fetchedData {
			t := convertCompetitionPointerToCompetitionPointer(d)
			result = append(result, t)
		}
	}
	return result
}

func convertInterfaceToCompetitionPointer(source interface{}) model.Competition {
	var i = source
	c, ok := i.(model.Competition)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.Competition", fmt.Sprintf("%T\n", source))
		return model.Competition{}
	}
	return c
}

func convertInterfaceToCompetitionDetailPointer(source interface{}) model.CompetitionDetail {
	var i = source
	c, ok := i.(model.CompetitionDetail)
	if !ok {
		log.Printf("Unable to convert type (%s) to model.CompetitionDetail", fmt.Sprintf("%T\n", source))
		return model.CompetitionDetail{}
	}
	return c
}

func convertCompetitionPointerToCompetitionPointer(fetchedData model.Competition) dto.Competition {
	if !common.IsStructEmpty(fetchedData) {
		return dto.Competition{
			ID:   fetchedData.ID,
			Name: fetchedData.Name,
			Year: fetchedData.Year,
			Type: fetchedData.ClimbingTypes[0],
		}
	}
	return dto.Competition{}
}

func convertSectionResultAndSectionMapToSectionDto(sr model.SectionResult, s map[string]model.Section, climbingType model.ClimbingType) dto.Section {
	ss := "?"
	for s2 := range s {
		so := s[s2]
		if so.ID == sr.Section {
			ss = so.Name
		}
	}
	var section dto.Section
	if climbingType == model.Boulder {
		section = dto.Section{
			Name:      ss,
			Tops:      sr.Tops,
			Zones:     sr.Zones,
			TopTries:  sr.TopTries,
			ZoneTries: sr.ZoneTries,
		}
	} else {
		section = dto.Section{
			Name:      ss,
			Points:    sr.Points,
			Tops:      sr.Tops,
			Zones:     sr.Zones,
			TopTries:  sr.TopTries,
			ZoneTries: sr.ZoneTries,
		}
	}
	return section
}
