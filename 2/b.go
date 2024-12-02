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

func B() {
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
		if validateReport(report, true) {
			fmt.Printf("%d is valid\n", i)
			valid++
		}

	}

	fmt.Printf("Valid: %d\n", valid)

}

func validateReport(report []int, retry bool) bool {
	isAsc := report[0] > report[1]
	for i := 0; i < len(report)-1; i++ {

		isVAsc := report[i] > report[i+1]

		isSame := report[i] == report[i+1]
		isDifferent := isAsc != isVAsc
		isDiff := absInt(report[i], report[i+1]) > 3

		if isSame || isDifferent || isDiff {
			if !retry {
				return false
			}
			if i != 0 {
				if validateReport(copyRemoveIdx(report, i-1), false) {
					return true
				}
			}

			if validateReport(copyRemoveIdx(report, i), false) {
				return true
			}
			if validateReport(copyRemoveIdx(report, i+1), false) {
				return true
			}
			return false
		}
	}

	return true
}

func copyRemoveIdx(data []int, idx int) []int {
	res := make([]int, 0)
	res = append(res, data[:idx]...)
	res = append(res, data[idx+1:]...)
	return res
}
