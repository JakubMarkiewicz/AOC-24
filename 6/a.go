package main

import (
	"fmt"
	"strings"
	"time"
)

// var Yellow = "\033[229m"
var Yellow = fmt.Sprintf("\u001B[48;2;%d;%d;%dm", 255, 0, 0)
var Reset = "\033[0m"

func A(input [][]string, startX, startY int) {
	visualizationCopy := deepCopy(input)
	visLast := [][]int{}

	fmt.Println(visualizationCopy)

	x := startX
	y := startY

	// 0 up, 1 right, 2 down, 3 left
	dir := 0

	visited := make(map[string]bool)

	for inBound(x, y, len(input[0]), len(input)) {
		if input[y][x] == "#" {
			switch dir {
			case 0:
				y++
				dir = 1
			case 1:
				x--
				dir = 2
			case 2:
				y--
				dir = 3
			case 3:
				x++
				dir = 0
			}
		} else {
			visLast = append(visLast, []int{y, x})

			for i := 0; i < 150 && i < len(visLast); i++ {
				v := visLast[len(visLast)-1-i]
				visualizationCopy[v[0]][v[1]] = fmt.Sprintf("\u001B[48;2;%d;%d;%dm.%v", 255-i, 0, 0, Reset)
			}

			visited[string(x)+string(y)] = true

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
		}
		printCurrState(visualizationCopy)

	}

	fmt.Println(len(visited))

}
func printCurrState(input [][]string) {
	var builder strings.Builder
	for _, i := range input {
		builder.WriteString(strings.Join(i, ""))
		builder.WriteRune('\n')
	}

	fmt.Print("\033[H\033[2J")
	fmt.Print(builder.String())
	time.Sleep(2 * time.Millisecond)

}

func deepCopy(val [][]string) [][]string {
	v := make([][]string, len(val))

	for j, c := range val {
		row := make([]string, len(val[0]))
		for i, v := range c {
			row[i] = v
		}
		v[j] = row
	}
	return v
}

func inBound(x, y, xSize, ySize int) bool {
	return x >= 0 && x < xSize && y >= 0 && y < ySize
}
