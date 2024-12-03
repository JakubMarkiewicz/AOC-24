package main

import (
	"fmt"
	"log"
	"strconv"
)

const (
	dont = "don't()"
	do   = "do()"
)

func B(sample bool) {
	lines, err := getInput(sample)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	linesRes := make([]int, len(lines))

	stage := 0
	for iL, line := range lines {
		// 0 - looking for mul, 1 - left, 2 - right, 3 - don't
		var l, r string

		for i := 0; i < len(line); i++ {
			c := line[i]

			if i < len(line)-len(dont) && line[i:i+len(dont)] == dont {
				stage = 3
				i += len(dont) - 1
			}

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
			case 3:
				if i < len(line)-len(do) && line[i:i+len(do)] == do {
					stage = 0
					i += len(do) - 1
				}

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
