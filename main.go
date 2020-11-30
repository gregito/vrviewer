package main

import (
	"fmt"
	"github.com/gregito/vrviewer/comp"
	"log"
	"strconv"
)

func main() {
	competition, err := comp.GetCompetitorOnCompetitionByCompetitionIdAndCompetitorName(33, "Farkas Tamás")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Name: " + competition.Name)
	fmt.Println()
	fmt.Println("Position: " + competition.CurrentPosition)
	fmt.Println()
	for _, r := range competition.SectionResults {
		fmt.Println(r.Name + ":")
		fmt.Println("Tops: \t\t" + strconv.FormatInt(r.Tops, 10))
		fmt.Println("Zones: \t\t" + strconv.FormatInt(r.Zones, 10))
		fmt.Println("Top tries: \t" + strconv.FormatInt(r.TopTries, 10))
		fmt.Println("Zone tries: \t" + strconv.FormatInt(r.ZoneTries, 10))
		fmt.Println()
	}
}
