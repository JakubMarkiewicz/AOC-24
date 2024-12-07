package main

import "fmt"

func A(input []I) {
	res := 0
	for _, v := range input {
		if isValid(v.vals[1:], v.vals[0], v.target) {
			res += v.target
		}
	}

	fmt.Println(res)

}

func isValid(input []int, curr, target int) bool {

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

	return isValid(iC, curr*f, target) || isValid(iC, curr+f, target)

}
