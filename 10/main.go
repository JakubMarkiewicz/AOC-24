package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	next := flag.Bool("next", false, "is second part")
	sample := flag.Bool("sample", false, "is sample data")

	flag.Parse()

	lines := parseInput(*sample)

	if *next {
		B(lines)
	} else {
		A(lines)
	}

	fmt.Println(time.Since(start))
}

func parseInput(sample bool) [][]int {
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

	var res [][]int

	for sc.Scan() {
		line := sc.Text()

		lineI := make([]int, len(line))

		for i, v := range line {
			vI, _ := strconv.Atoi(string(v))
			lineI[i] = vI
		}
		res = append(res, lineI)
	}

	return res
}
