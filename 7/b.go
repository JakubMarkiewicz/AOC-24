package main

import (
	"fmt"
	"strconv"
	"sync"
)

func B(input []I) {
	sums := make(chan int, len(input))
	var wg sync.WaitGroup

	for _, v := range input {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if isValidB(v.vals[1:], v.vals[0], v.target) {
				sums <- v.target
			} else {
				sums <- 0
			}
		}()
	}

	wg.Wait()
	close(sums)

	res := 0
	for x := range sums {
		res += x
	}

	fmt.Println(res)

}

func isValidB(input []int, curr, target int) bool {

	if len(input) == 0 {
		return curr == target
	}

	if curr > target {
		return false
	}

	f := input[0]
	iC := make([]int, len(input)-1)

	for i := 1; i < len(input); i++ {
		iC[i-1] = input[i]
	}

	return isValidB(iC, curr*f, target) || isValidB(iC, curr+f, target) || isValidB(iC, concatInt(curr, f), target)
}

func concatInt(a int, b int) int {
	s1 := strconv.Itoa(a)
	s2 := strconv.Itoa(b)

	v, _ := strconv.Atoi(s1 + s2)

	return v
}
