package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
)

func main() {
	next := flag.Bool("next", false, "Is second part")
	flag.Parse()

	lines := parseInput()

	if *next {
		fmt.Println(B(lines))
	} else {
		fmt.Println(A(lines))
	}
}

func parseInput() []string {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}

	sc := bufio.NewScanner(file)

	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
