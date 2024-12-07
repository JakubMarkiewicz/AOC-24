package main

import (
	"bufio"
	"flag"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	next := flag.Bool("next", false, "is second part")
	sample := flag.Bool("sample", false, "is sample data")

	flag.Parse()

	input, startX, startY := parseInput(*sample)

	if *next {
		B(input, startX, startY)
	} else {
		A(input, startX, startY)
	}
}

func parseInput(sample bool) ([][]string, int, int) {
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

	lines := [][]string{}
	x := -1
	y := -1

	sc := bufio.NewScanner(file)

	i := 0
	for sc.Scan() {
		line := sc.Text()

		lines = append(lines, strings.Split(line, ""))

		if x == -1 {
			for j, c := range line {
				if c == '^' {
					x = j
					y = i
				}
			}
		}

		i++

	}

	return lines, x, y
}
