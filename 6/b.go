package main

import (
	"fmt"
)

func B(input [][]string, startX, startY int) {
	x, y := startX, startY

	// 0 up, 1 right, 2 down, 3 left
	dir := 0

	loops := 0

	visited := make(map[string]bool)

	for inBound(x, y, len(input[0]), len(input)) {
		if input[y][x] == "#" {
			x, y, dir = changeDir(x, y, dir)
		} else {
			_, prs := visited[string(x)+string(y)]
			if checkLoop(input, x, y, dir) && !prs {
				loops++
			}
			visited[string(x)+string(y)] = true
			x, y = move(x, y, dir)
		}

	}

	fmt.Println(loops)

}

func checkLoop(input [][]string, x, y, dir int) bool {
	xC, yC := x, y

	x, y = moveBack(x, y, dir)
	dir = (dir + 1) % 4

	visitedHash := make(map[string]bool)
	for inBound(x, y, len(input[0]), len(input)) {
		id := fmt.Sprintf("%v|%v|%v", x, y, dir)

		e := visitedHash[id]
		if e {
			return true
		}
		visitedHash[id] = true
		if input[y][x] == "#" || (y == yC && x == xC) {
			x, y, dir = changeDir(x, y, dir)
		} else {
			x, y = move(x, y, dir)
		}

	}

	return false

}

func move(x, y, dir int) (int, int) {
	switch dir {
	case 0:
		y--
	case 1:
		x++
	case 2:
		y++
	case 3:
		x--
	}

	return x, y
}

func moveBack(x, y, dir int) (int, int) {
	switch dir {
	case 0:
		y++
	case 1:
		x--
	case 2:
		y--
	case 3:
		x++
	}
	return x, y
}

func changeDir(x, y, dir int) (int, int, int) {
	switch dir {
	case 0:
		y++
	case 1:
		x--
	case 2:
		y--
	case 3:
		x++
	}

	dir = (dir + 1) % 4

	return x, y, dir

}
