package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
)

func A(sample bool) {
	lines, err := getInput(sample)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	linesRes := make([]int, len(lines))

	for iL, line := range lines {
		// 0 - looking for mul, 1 - left, 2 - right
		stage := 0
		var l, r string

		for i := 0; i < len(line); i++ {
			c := line[i]

			switch stage {
			case 0:
				if i < len(line)-5 && line[i:i+3] == "mul" && line[i+3] == '(' {
					stage = 1
					l = ""
					r = ""
					i += 3
				}
				continue
			case 1:
				if c == ',' {
					stage = 2
					continue
				}
				if !isCharNumeric(c) {
					stage = 0
				}
				l += string(c)
				continue
			case 2:
				if c == ')' {
					stage = 0
					lN, errLeft := strconv.Atoi(l)
					rN, errRight := strconv.Atoi(r)
					if errLeft != nil || errRight != nil {
						stage = 0
						continue
					}
					linesRes[iL] += lN * rN
					stage = 0
				}
				if !isCharNumeric(c) {
					stage = 0
				}
				r += string(c)
			}
			continue
		}
	}

	sum := 0
	for _, v := range linesRes {
		sum += v
	}

	fmt.Println(sum)
}

func getInput(sample bool) ([]string, error) {
	var file *os.File
	var err error
	if sample {
		file, err = os.OpenFile("./sample.txt", os.O_RDONLY, fs.ModePerm)
	} else {
		file, err = os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)
	}

	if err != nil {
		return nil, err
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	lines := []string{}

	for sc.Scan() {
		lines = append(lines, sc.Text())
	}

	return lines, nil

}

func isCharNumeric(c byte) bool {
	return c >= 48 && c <= 57

}
