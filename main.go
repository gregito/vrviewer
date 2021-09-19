package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gregito/vrviewer/comp"
	"github.com/gregito/vrviewer/comp/dto"
	"github.com/gregito/vrviewer/comp/log"
	"github.com/gregito/vrviewer/comp/metrics"
	"github.com/gregito/vrviewer/comp/model"

	"github.com/lensesio/tableprinter"
)

var competitionResults []model.CompetitionDetail
var singleFetchDurations []time.Duration
var totalFetchTime time.Duration
var compType model.ClimbingType
var competitionName string
var names []string
var year int64
var args []string
var competitionListingRequested bool

func init() {
	args = os.Args[1:]
	fmt.Println()
	if args != nil && len(args) != 0 {
		logInputArgs(args)
		logEnvs()
		findAndParseInputs()
		fetchData()
	} else {
		fmt.Println("No competitor name has provided.")
		os.Exit(2)
	}
}

func main() {
	listContent()
	metrics.ShowMeasurementsIfHaveAny(singleFetchDurations, totalFetchTime)
}

func fetchData() {
	fmt.Println("Fetching competition data...")
	start := time.Now()
	defer func() {
		totalFetchTime = time.Since(start)
	}()
	competitionResults, singleFetchDurations = comp.ListAllCompetitionDetail(competitionName, year, compType)
	fmt.Println("\nFetching done.")
}

func listContent() {
	fmt.Printf("\n")
	if competitionListingRequested {
		competitions := collectCompetitions()
		if len(competitions) == 0 {
			fmt.Printf("No competition has found for year: %d\n\n", year)
		} else {
			tableprinter.Print(os.Stdout, competitions)
			fmt.Println()
		}
	} else {
		for _, name := range names {
			competitorResults := comp.GetCompetitorResults(name, competitionResults)
			if competitorResults != nil && len(competitorResults) > 0 {
				fmt.Println("------------------")
				fmt.Println("Name: " + name)
				for _, result := range competitorResults {
					fmt.Println("==================")
					if len(competitionName) > 0 {
						fmt.Println()
					} else {
						fmt.Printf("%s - %s\n\n", result.CompetitionName, result.Category)
					}
					for _, ageGroup := range result.AgeGroupResult {
						printAgeGroup(result.CompetitionFinished, ageGroup)
					}
				}
			} else {
				fmt.Println("No competitor has been found with name: " + name)
			}
		}
	}
}

func collectCompetitions() []dto.BasicCompetition {
	var list []dto.BasicCompetition
	sort.SliceStable(competitionResults, func(i, j int) bool {
		return competitionResults[i].ID < competitionResults[j].ID
	})
	for _, compResults := range competitionResults {
		var comp dto.BasicCompetition
		if len(string(compType)) != 0 {
			comp = dto.BasicCompetition{
				Name:   compResults.Name,
				Type:   compType,
				Status: compResults.Status,
			}
		} else {
			if len(compResults.Partitions) >= 1 {
				comp = dto.BasicCompetition{
					Name:   compResults.Name,
					Type:   compResults.Partitions[0].ClimbingType,
					Status: compResults.Status,
				}
			} else {
				comp = dto.BasicCompetition{
					Name:   compResults.Name,
					Type:   "Unknown",
					Status: compResults.Status,
				}
			}
		}
		list = append(list, comp)
	}
	return list
}

func printAgeGroup(isCompetitionFinished bool, ag dto.AgeGroupResult) {
	fmt.Println("Group: " + ag.AgeGroup)
	printPosition(isCompetitionFinished, ag.FinalPosition)
	printSections(ag.Results)
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
			fmt.Println()
		}
		fmt.Println()
	}
}

func printPosition(isCompetitionFinished bool, position int64) {
	ps := " position"
	if isCompetitionFinished {
		ps = "Final" + ps
	} else {
		ps = "Current" + ps
	}
	fmt.Printf("%s: %d\n", ps, position)
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

func findAndParseInputs() {
	for _, arg := range args {
		checkCompetitionListingRequested(arg)
		findDesiredYearInArgs(arg)
		findCompetitionTypeInArgs(arg)
		findNamesInArgs(arg)
		findDesiredCompetitionByNameInArgs(arg)
	}
	if competitionListingRequested {
		if year == 0 {
			fmt.Println("If competition listing requested then a proper year has to be given!")
			os.Exit(2)
		}
	}
}

func checkCompetitionListingRequested(arg string) {
	if strings.HasPrefix(arg, "--list-competitions") {
		competitionListingRequested = true
		log.Println("Competition listing requested.")
	}
}

func findNamesInArgs(arg string) {
	if strings.HasPrefix(arg, "--names") {
		splitArg := splitByEqualSign(arg)
		if len(splitArg) == 2 {
			rawInputNames := strings.Split(splitArg[1], ",")
			for _, name := range rawInputNames {
				names = append(names, strings.TrimSpace(name))
			}
			log.Printf("The given name(s) has been provided by the user: %s\n", names)
		} else {
			names = make([]string, 0)
			log.Println()
		}
	}
}

func findDesiredYearInArgs(arg string) {
	if strings.HasPrefix(arg, "--year") {
		splitArg := splitByEqualSign(arg)
		var err error
		year, err = strconv.ParseInt(splitArg[1], 10, 64)
		if err != nil {
			log.Printf("Unable to convert string year into int64 (%s -> int64)\n", splitArg[0])
			year = comp.ValueToDisableYearFilter
		} else {
			log.Println("Year filter value has been set to: " + splitArg[1])
		}
	}
}

func findDesiredCompetitionByNameInArgs(arg string) {
	if strings.HasPrefix(arg, "--competition-name") {
		splitArg := splitByEqualSign(arg)
		if len(splitArg) == 2 && len(splitArg[1]) > 0 {
			competitionName = splitArg[1]
		} else {
			log.Println("Unable to set competition name")
		}

	}
}

func findCompetitionTypeInArgs(arg string) {
	if strings.HasPrefix(arg, "--type") {
		splitArg := splitByEqualSign(arg)
		if len(splitArg) == 2 {
			if splitArg[1] == string(model.Lead) {
				compType = model.Lead
				log.Println("Competition type has been set to: " + string(model.Lead))
			} else if splitArg[1] == string(model.Boulder) {
				compType = model.Boulder
				log.Println("Competition type has been set to: " + string(model.Boulder))
			} else {
				log.Println("Unable to identify competition type because the user has not provided a valid one. We are going to move on without filtering this")
			}
		}
	}
}

func splitByEqualSign(s string) []string {
	return strings.Split(s, "=")
}
