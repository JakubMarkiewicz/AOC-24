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

	input := parseInput(*sample)

	if *next {
		B(input)
	} else {
		A(input)
	}

	fmt.Println(time.Since(start))
}

type I struct {
	target int
	vals   []int
}

func parseInput(sample bool) []I {
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

	res := []I{}

	for sc.Scan() {
		s := strings.Split(sc.Text(), ":")
		target, _ := strconv.Atoi(s[0])
		s2 := strings.Split(strings.TrimSpace(s[1]), " ")
		s2I := make([]int, len(s2))

		for i := 0; i < len(s2); i++ {
			v, _ := strconv.Atoi(s2[i])
			s2I[i] = v
		}

		res = append(res, I{target: target, vals: s2I})

	}

	return res

}
