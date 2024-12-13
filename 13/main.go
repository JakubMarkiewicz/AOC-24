package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	next := flag.Bool("next", false, "Is next part")

	flag.Parse()

	games := getInput(parseInput())

	if *next {
		fmt.Println(B(games))
	} else {
		fmt.Println(A(games))
	}
}

type Game struct {
	A []int
	B []int
	P []int
}

func getInput(str string) []Game {
	r, _ := regexp.Compile("(\\d{1,})")
	games := strings.Split(str, "\n\n")

	var res []Game

	for _, game := range games {
		g := Game{}
		lines := strings.Split(game, "\n")
		for i := 0; i < 3; i++ {
			matches := r.FindAllString(lines[i], -1)
			var ints []int
			for _, match := range matches {
				i, _ := strconv.Atoi(match)
				ints = append(ints, i)
			}

			if i == 0 {
				g.A = ints
			} else if i == 1 {
				g.B = ints
			} else {
				g.P = ints
			}
		}
		res = append(res, g)
	}

	return res

}

func parseInput() string {
	file, err := os.ReadFile("./input.txt")

	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}

	return string(file)
}
