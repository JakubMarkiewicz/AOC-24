package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	next := flag.Bool("next", false, "Is second part")
	flag.Parse()

	data := parseInput(parseFile())
	if *next {
		B(data, 103, 101)
	} else {
		fmt.Println(A(data, 100, 103, 101))
	}

}

type Coord struct {
	x, y int
}

type Robot struct {
	pos  Coord
	move Coord
}

func parseInput(lines []string) []Robot {
	var res []Robot

	for _, line := range lines {
		rxp, _ := regexp.Compile("(-?\\d{1,})")
		data := rxp.FindAllString(line, -1)

		if len(data) != 4 {
			log.Fatalf("Invalid entry %v", line)
		}

		dataInts := make([]int, 4)

		for i, v := range data {
			vInt, _ := strconv.Atoi(v)
			dataInts[i] = vInt
		}

		res = append(res, Robot{
			Coord{dataInts[0], dataInts[1]},
			Coord{dataInts[2], dataInts[3]},
		})
	}

	return res

}

func parseFile() []string {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file; %v", err)
	}

	sc := bufio.NewScanner(file)

	var lines []string

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines

}
