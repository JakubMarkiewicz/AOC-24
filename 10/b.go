package main

import (
	"fmt"
)

func B(input [][]int) {
	res := 0

	for y, row := range input {
		for x, v := range row {
			if v == 0 {
				res += calculatePathsB(y, x, input, [][]int{})
			}
		}

	}

	fmt.Println(res)

}

func calculatePathsB(y, x int, input [][]int, visited [][]int) int {
	if hasB(visited, []int{y, x}) {
		return 0
	}
	visited = append(visited, []int{y, x})

	if input[y][x] == 9 {
		return 1
	}

	width := len(input[0])
	height := len(input)

	res := 0

	next := [][]int{}

	if y-1 >= 0 {
		next = append(next, []int{y - 1, x})
	}
	if y+1 < height {
		next = append(next, []int{y + 1, x})
	}
	if x-1 >= 0 {
		next = append(next, []int{y, x - 1})
	}
	if x+1 < width {
		next = append(next, []int{y, x + 1})
	}

	nextFiltered := [][]int{}

	for _, v := range next {
		if !hasB(visited, v) && input[v[0]][v[1]]-input[y][x] == 1 {
			nextFiltered = append(nextFiltered, v)
		}
	}

	for _, v := range nextFiltered {
		res += calculatePathsB(v[0], v[1], input, visited)
	}

	return res
}

func hasB(visited [][]int, v []int) bool {
	for _, vi := range visited {
		if vi[0] == v[0] && vi[1] == v[1] {
			return true
		}
	}
	return false
}
