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

	line := parseInput(*sample)

	if *next {
		B(line)
	} else {
		A(line)
	}

	fmt.Println(time.Since(start))
}

func parseInput(sample bool) string {
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

	line := ""

	for sc.Scan() {
		line = sc.Text()

	}

	return line
}
