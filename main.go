package main

import (
	"fmt"
	"github.com/gregito/vrviewer/comp"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/metrics"
	"github.com/gregito/vrviewer/comp/model"
	"os"
	"strconv"
	"time"
)

var competitionResults []model.CompetitionDetail
var singleFetchDurations []time.Duration
var totalFetchTime time.Duration

func init() {
	fetchData()
}

func main() {
	args := os.Args[1:]
	listStuff(competitionResults, args)
	metrics.ShowMeasurements(singleFetchDurations, totalFetchTime)
}

func fetchData() {
	start := time.Now()
	defer func() {
		totalFetchTime = time.Since(start)
	}()
	competitions, dur := comp.ListAllCompetitionsSimplified()
	singleFetchDurations = append(singleFetchDurations, dur)
	for _, competition := range competitions {
		res, err, dur := comp.GetCompetitionResultsByCompetitionId(competition.ID)
		singleFetchDurations = append(singleFetchDurations, dur)
		if err == nil {
			competitionResults = append(competitionResults, res)
		}
	}
}

func listStuff(competitionResults []model.CompetitionDetail, names []string) {
	for i, name := range names {
		fmt.Println("Name: " + name)

		competitorResults := comp.GetCompetitorResults(name, competitionResults)

		for _, result := range competitorResults {
			fmt.Printf("%s - %s\n", result.CompetitionName, result.Type)
			printSections(result.SectionResults)
		}
		if i < len(names)-1 {
			fmt.Println("------------------")
		}
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
