package main

import (
	"fmt"
	"strconv"
)

func A(line string) {
	parsed := parseDiskMap(line)

	l := 0
	r := len(parsed) - 1

	for l < r {
		lVal := parsed[l]
		rVal := parsed[r]

		if lVal == "." {
			if rVal == "." {
				r--
				continue
			} else {
				parsed[l] = rVal
				parsed[r] = "."
				l++
				r--
			}
		} else {
			l++
		}
	}

	res := 0

	for i, v := range parsed {
		if v == "." {
			break
		}
		vInt, _ := strconv.Atoi(v)
		res += (i * vInt)
	}

	fmt.Println(res)
}

func parseDiskMap(line string) []string {
	res := []string{}

	id := 0

	for i := 0; i < len(line); i++ {
		v, _ := strconv.Atoi(string(line[i]))

		if i%2 == 0 {
			idS := strconv.Itoa(id)
			res = append(res, buildArr(idS, v)...)
			id++
		} else {
			res = append(res, buildArr(".", v)...)
		}
	}

	return res
}

func buildArr(v string, l int) []string {
	res := make([]string, l)

	for i := 0; i < l; i++ {
		res[i] = v
	}

	return res
}
