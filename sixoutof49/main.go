package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/XelaRellum/sixoutof49"
)

var (
	max   int
	count int
)

func init() {
	flag.IntVar(&max, "max", 49, "maximal number to use")
	flag.IntVar(&count, "count", 6, "number of throws")
}

func main() {
	flag.Parse()

	numbers, err := sixoutof49.Create(1, max, count)
	if err != nil {
		panic(err)
	}

	numbersStr := strings.Trim(fmt.Sprintf("%v", numbers), "[]")
	fmt.Printf("sixoutof49: your lucky numbers are: %v\n", numbersStr)
}
