package main

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp"
	"github.com/gregito/vrviewer/comp/model"
)

func main() {
	kind := model.Lead
	comp, err := comp.ListCompetitionsByKind(&kind)
	if err != nil {
		log.Fatalln(err)
	}
	for _, c := range comp {
		fmt.Printf("%d :: %s :: %s\n", c.Year, c.Name, c.Type)
	}
}
