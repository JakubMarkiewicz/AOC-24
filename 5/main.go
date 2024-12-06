package main

import (
	"bufio"
	"flag"
	"io/fs"
	"log"
	"os"
)

func main() {
	next := flag.Bool("next", false, "Is second part")
	sample := flag.Bool("sample", false, "Is sample data")

	flag.Parse()

	input := parseFile(*sample)

	if *next {
		B(input)
	} else {
		A(input)
	}
}

func parseFile(sample bool) []string {
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

	defer file.Close()

	sc := bufio.NewScanner(file)
	lines := []string{}

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines
}
