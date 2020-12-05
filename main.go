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
var invalidStart bool
var args []string

func init() {
	args = os.Args[1:]
	if args != nil && len(args) != 0 {
		fetchData()
		invalidStart = false
	} else {
		invalidStart = true
	}
}

func main() {
	if invalidStart {
		fmt.Println("No competitor name has provided.")
		os.Exit(2)
	} else {
		listStuff(competitionResults, args)
		metrics.ShowMeasurements(singleFetchDurations, totalFetchTime)
	}
}

func fetchData() {
	fmt.Println("Fetching competition data...")
	start := time.Now()
	defer func() {
		totalFetchTime = time.Since(start)
	}()
	fmt.Println("Fetching competitions")
	competitions, dur := comp.ListAllCompetitionsSimplified()
	singleFetchDurations = append(singleFetchDurations, dur)
	fmt.Println("About to collect all competition results. This could take a wile depending on your network bandwidth.")
	for i, competition := range competitions {
		getProgressBar(i, len(competitions))
		res, err, dur := comp.GetCompetitionResultsByCompetitionId(competition.ID)
		singleFetchDurations = append(singleFetchDurations, dur)
		if err == nil {
			competitionResults = append(competitionResults, res)
		}
	}
	fmt.Println("\nFetching done.")
}

func listStuff(competitionResults []model.CompetitionDetail, names []string) {
	fmt.Printf("\n")
	for i, name := range names {
		competitorResults := comp.GetCompetitorResults(name, competitionResults)
		if competitorResults != nil && len(competitorResults) > 0 {
			fmt.Println("Name: " + name)
			for _, result := range competitorResults {
				fmt.Printf("%s - %s\n", result.CompetitionName, result.Type)
				printSections(result.SectionResults)
			}
			if i < len(names)-1 {
				fmt.Println("------------------")
			}
		} else {
			fmt.Println("No competitor has been found with name: " + name)
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

func getProgressBar(curr int, size int) {
	pb := "["
	for i := 0; i < curr; i++ {
		pb = pb + "|"
	}
	for i := curr; i < size-1; i++ {
		pb = pb + "-"
	}
	pb = pb + "]"
	fmt.Printf("\r%s", pb)
}
