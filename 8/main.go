package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"time"
)

func main() {
	start := time.Now()
	next := flag.Bool("next", false, "is second part")
	sample := flag.Bool("sample", false, "is sample data")

	flag.Parse()

	lines, input, height, width := parseInput(*sample)

	if *next {
		B(lines, input, height, width)
	} else {
		A(lines, input, height, width)
	}

	fmt.Println(time.Since(start))
}

type Coord struct {
	y, x int
}

func parseInput(sample bool) ([]string, map[rune][]Coord, int, int) {
	var file *os.File
	var err error

	if sample {
		file, err = os.OpenFile("./sample.txt", os.O_RDONLY, fs.ModePerm)
	} else {
		file, err = os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)
	}

	if err != nil {
		log.Fatalf("Couldn't open file: %v", err)
	}

	sc := bufio.NewScanner(file)

	res := make(map[rune][]Coord)

	lines := []string{}

	y := 0
	width := 0
	for sc.Scan() {
		line := sc.Text()
		lines = append(lines, line)
		width = len(line)

		for x, v := range line {
			if v != '.' {
				_, prs := res[v]
				if prs {
					res[v] = append(res[v], Coord{y, x})
				} else {
					res[v] = []Coord{{y, x}}
				}
			}

		}
		y++

	}

	return lines, res, y, width
}
