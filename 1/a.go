package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func A() {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var l []int
	var r []int

	for scanner.Scan() {
		txt := strings.Split(scanner.Text(), "   ")

		lVal, err := strconv.Atoi(txt[0])
		if err != nil {
			log.Fatalf("Couldn't convert %q to int", lVal)
		}

		rVal, err := strconv.Atoi(txt[1])
		if err != nil {
			log.Fatalf("Couldn't convert %q to int", rVal)
		}

		l = append(l, lVal)
		r = append(r, rVal)
	}

	sort.Ints(l)
	sort.Ints(r)

	res := 0

	for i := 0; i < len(l); i++ {
		diff := r[i] - l[i]
		if diff < 0 {
			res += (diff * -1)
		} else {
			res += diff
		}
	}

	fmt.Println(res)
}
