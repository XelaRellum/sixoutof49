package main

import (
	"fmt"
	"strings"

	"github.com/XelaRellum/sixoutof49"
)

func main() {
	numbers, err := sixoutof49.Create()
	if err != nil {
		panic(err)
	}

	numbersStr := strings.Trim(fmt.Sprintf("%v", numbers), "[]")
	fmt.Printf("sixoutof49: your lucky numbers are: %v\n", numbersStr)
}
