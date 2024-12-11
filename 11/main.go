package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	next := flag.Bool("next", false, "is second part")
	sample := flag.Bool("sample", false, "is sample data")

	flag.Parse()

	stones := parseInput(*sample)

	if *next {
		B(stones)
	} else {
		A(stones)
	}

	fmt.Println(time.Since(start))
}

func parseInput(sample bool) []int {
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

	var res []int

	for sc.Scan() {
		line := sc.Text()

		for _, v := range strings.Split(line, " ") {
			vI, _ := strconv.Atoi(string(v))
			res = append(res, vI)
		}
	}

	return res
}
