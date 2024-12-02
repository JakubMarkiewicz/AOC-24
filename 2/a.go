package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strconv"
	"strings"
)

func A() {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)

	if err != nil {
		log.Fatalf("Couldn't read file: %v", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	reports := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()

		var report []int

		for _, v := range strings.Split(line, " ") {
			i, err := strconv.Atoi(v)

			if err != nil {
				log.Fatalf("Couldn't convert to int: %v", v)
			}

			report = append(report, i)
		}
		reports = append(reports, report)
	}

	valid := 0

	for i, report := range reports {
		isValid := true
		isAsc := report[1] > report[0]
		for i := 1; i < len(report); i++ {

			isVAsc := report[i] > report[i-1]

			isSame := report[i] == report[i-1]
			isDifferent := isAsc != isVAsc
			isDiff := absInt(report[i], report[i-1]) > 3

			if isSame || isDifferent || isDiff {
				isValid = false
				break
			}

		}

		if isValid {
			fmt.Printf("Raport %d is save\n", i)
			valid++
		}

	}

}

func absInt(x int, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
