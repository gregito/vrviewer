package main

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp"
)

func main() {
	comp, err := comp.GetCompetitionResultsByCompetitionId(33)
	if err != nil {
		log.Fatalln(err)
	}
	for _, c := range comp.Partitions {
		for _, p := range c.Results {
			if len(p.SectionResults) > 1 {
				fmt.Printf("%s-%s :: %s (%d.)\n", c.Gender, c.AgeGroup, p.Name, p.Position)
			}
		}
	}
}
