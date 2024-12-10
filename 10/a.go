package main

import (
	"fmt"
)

func A(input [][]int) {
	res := 0

	for y, row := range input {
		for x, v := range row {
			if v == 0 {
				res += calculatePathsA(y, x, input)
			}
		}

	}

	fmt.Println(res)

}

func calculatePathsA(y, x int, input [][]int) int {
	width := len(input[0])
	height := len(input)

	visited := [][]int{}
	bfs := [][]int{{y, x}}

	res := 0

	for len(bfs) > 0 {

		first := bfs[0]
		y, x := first[0], first[1]
		bfs = bfs[1:]

		if hasA(visited, first) {
			continue
		}
		visited = append(visited, first)

		if input[y][x] == 9 {
			res++
			continue
		}

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

		for _, v := range next {
			if !hasA(visited, v) && input[v[0]][v[1]]-input[y][x] == 1 {
				bfs = append(bfs, v)
			}
		}
	}

	return res
}

func hasA(visited [][]int, v []int) bool {
	for _, vi := range visited {
		if vi[0] == v[0] && vi[1] == v[1] {
			return true
		}
	}
	return false
}
