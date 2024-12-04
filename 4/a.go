package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

const word = "XMAS"

func A(sample bool) {
	input := parseInput(sample)

	res := 0

	for rI, r := range input {
		for cI := range r {
			res += count(rI, cI, input)
		}

	}

	fmt.Println(res)

}

func count(rI, cI int, data [][]string) int {
	height := len(data)
	length := len(data[0])

	if data[rI][cI] != "X" {
		return 0
	}

	check := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	res := 0

	for _, v := range check {
		temp := "X"
		r := rI + v[0]
		c := cI + v[1]
		for {
			if temp == word {
				res++
				break
			}
			if r < 0 || r >= height {
				break
			}
			if c < 0 || c >= length {
				break
			}
			dV := data[r][c]
			next := string(word[len(temp)])
			if dV != next {
				break
			}
			temp = temp + dV
			r += v[0]
			c += v[1]
		}

	}

	return res

}

func parseInput(sample bool) [][]string {
	var file *os.File
	var err error

	if sample {
		file, err = os.OpenFile("./sample.txt", os.O_RDONLY, fs.ModePerm)
	} else {
		file, err = os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)
	}

	if err != nil {
		log.Fatalf("Error parsing input: %v", err)
	}

	defer file.Close()

	res := [][]string{}

	sc := bufio.NewScanner(file)

	for sc.Scan() {
		res = append(res, strings.Split(sc.Text(), ""))
	}

	return res
}
