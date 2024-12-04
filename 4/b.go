package main

import (
	"fmt"
)

func B(sample bool) {
	input := parseInput(sample)

	res := 0

	for rI := 1; rI < len(input)-1; rI++ {
		for cI := 1; cI < len(input[0])-1; cI++ {
			if check(rI, cI, input) {
				res++
			}
		}
	}

	fmt.Println(res)

}

func check(rI, cI int, data [][]string) bool {

	if data[rI][cI] != "A" {
		return false
	}

	l := data[rI-1][cI-1] + data[rI+1][cI+1]
	r := data[rI+1][cI-1] + data[rI-1][cI+1]

	return (l == "MS" || l == "SM") && (r == "MS" || r == "SM")
}
