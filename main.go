package main

import (
	"fmt"
	"github.com/gregito/vrviewer/comp"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
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
		logInputArgs(args)
		logEnvs()
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
		listContent(args)
		metrics.ShowMeasurementsIfHaveAny(singleFetchDurations, totalFetchTime)
	}
}

func fetchData() {
	fmt.Println("Fetching competition data...")
	start := time.Now()
	defer func() {
		totalFetchTime = time.Since(start)
	}()
	competitionResults, singleFetchDurations = comp.ListAllCompetitionDetail()
	fmt.Println("\nFetching done.")
}

func listContent(names []string) {
	fmt.Printf("\n")
	for i, name := range names {
		competitorResults := comp.GetCompetitorResults(name, competitionResults)
		if competitorResults != nil && len(competitorResults) > 0 {
			fmt.Println("Name: " + name)
			for _, result := range competitorResults {
				fmt.Println("------------------")
				fmt.Printf("%s - %s\n", result.CompetitionName, result.Type)
				printPosition(result)
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
		fmt.Println("\n" + r.Name + ":")
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

func logInputArgs(args []string) {
	out := ""
	for i := 0; i < len(args)-1; i++ {
		out = out + "\"" + args[i] + "\", "
	}
	out = out + "\"" + args[len(args)-1] + "\""
	log.Printf("Program arguments: %s", out)
}

func logEnvs() {
	log.Printf("Environment variables: %s", os.Environ())
}

func printPosition(result dto.CompetitorResult) {
	ps := " position"
	if result.CompetitionFinished {
		ps = "Final" + ps
	} else {
		ps = "Current" + ps
	}
	fmt.Printf("%s: %s\n", ps, result.CurrentPosition)
}
