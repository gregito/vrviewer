package main

import (
	"fmt"
	"github.com/gregito/vrviewer/comp"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/model"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	defer func() {
		fmt.Println("Execution Time: ", time.Since(start))
	}()

	competitions := comp.ListAllCompetitionsSimplified()

	var competitionResults []model.CompetitionDetail
	for _, competition := range competitions {
		res, err := comp.GetCompetitionResultsByCompetitionId(competition.ID)
		if err == nil {
			competitionResults = append(competitionResults, res)
		}
	}
	listStuff(competitionResults, "Mészáros Gergely")
}

func listStuff(competitionResults []model.CompetitionDetail, names ...string) {
	for _, name := range names {
		fmt.Println("Name: " + name)

		competitorResults := comp.GetCompetitorResults(name, competitionResults)

		for _, result := range competitorResults {
			fmt.Printf("%s - %s\n", result.CompetitionName, result.Type)
			printSections(result.SectionResults)
		}
		fmt.Println("------------------")
	}

}

func printSections(sr []dto.Section) {
	for _, r := range sr {
		fmt.Println(r.Name + ":")
		if r.Points > 0 {
			// in case of lead climbing competitions only the points that matters therefore all other fields should be empty
			fmt.Println("Points: \t" + strconv.FormatInt(r.Points, 10))
		} else {
			// on the other hand, on boulder competitions the tops, zones and their tries matters so printing point
			// is not necessary since it is always zero
			fmt.Println("Tops: \t\t" + strconv.FormatInt(r.Tops, 10))
			fmt.Println("Zones: \t\t" + strconv.FormatInt(r.Zones, 10))
			fmt.Println("Top tries: \t" + strconv.FormatInt(r.TopTries, 10))
			fmt.Println("Zone tries: \t" + strconv.FormatInt(r.ZoneTries, 10))
		}
		fmt.Println()
	}
}
