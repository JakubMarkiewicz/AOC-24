package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func B() {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var l []int
	countMap := make(map[int]int)

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
		_, exists := countMap[rVal]
		if !exists {
			countMap[rVal] = 1
		} else {
			countMap[rVal] += 1
		}
	}

	res := 0

	for _, v := range l {
		c, exists := countMap[v]
		if exists {
			res += (v * c)
		}
	}

	fmt.Println(res)
}
