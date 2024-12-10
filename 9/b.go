package main

import (
	"fmt"
	"slices"
	"strconv"
)

func B(line string) {
	parsed := parseDiskMapB(line)

	for r := len(parsed) - 1; r >= 0; r-- {
		rVal := parsed[r]

		if rVal.val == "." {
			continue
		}

		for l := 0; l < r; l++ {
			lVal := parsed[l]

			if lVal.val != "." {
				continue
			}

			diff := lVal.length - rVal.length

			if diff >= 0 {
				if diff == 0 {
					parsed[l].val = rVal.val
					parsed[r].val = "."
				} else {
					parsed = slices.Insert(parsed, l, rVal)
					parsed[l+1].length = diff
					parsed[r+1].val = "."
				}

				r = len(parsed) - 1
				break
			}
		}
	}

	i := 0
	res := 0

	for _, v := range parsed {
		if v.val == "." {
			i += v.length
			continue
		}
		vVal, _ := strconv.Atoi(v.val)
		for j := 0; j < v.length; j++ {
			res += (i * vVal)
			i++
		}
	}

	fmt.Println(res)
}

type D struct {
	val    string
	length int
}

func parseDiskMapB(line string) []D {
	res := []D{}

	id := 0

	for i := 0; i < len(line); i++ {
		v, _ := strconv.Atoi(string(line[i]))

		if v == 0 {
			continue
		}

		if i%2 == 0 {
			idS := strconv.Itoa(id)
			res = append(res, D{val: idS, length: v})
			id++
		} else {
			res = append(res, D{val: ".", length: v})
		}
	}

	return res
}
