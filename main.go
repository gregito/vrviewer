package main

import (
	"fmt"
	"log"

	"github.com/gregito/vrviewer/comp"
)

func main() {
	comp, err := comp.ListCompetitions(2020, comp.Boulder)
	if err != nil {
		log.Fatalln(err)
	}
	for _, c := range comp {
		fmt.Printf("%+v\n", c)
	}
}
