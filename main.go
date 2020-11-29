package main

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp"
)

func main() {
	kind := comp.Boulder
	comp, err := comp.ListCompetitionsByKind(&kind)
	if err != nil {
		log.Fatalln(err)
	}
	for _, c := range comp {
		fmt.Printf("%d :: %s :: %s\n", c.Year, c.Name, c.Type)
	}
}
