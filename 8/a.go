package main

import (
	"fmt"
	"strings"
)

func A(lines []string, input map[rune][]Coord, height, width int) {
	antiNodes := make(map[Coord]bool)

	for _, crs := range input {
		calcPerms(crs, antiNodes, height, width)
	}

	printRes(lines, antiNodes)
	fmt.Println(len(antiNodes))

}

func calcPerms(crds []Coord, aNs map[Coord]bool, height, width int) {
	f := crds[0]

	for _, s := range crds[1:] {
		yDif := f.y - s.y
		xDif := f.x - s.x

		c := Coord{y: f.y + yDif, x: f.x + xDif}

		if inbound(c, width, height) {
			aNs[c] = true
		}

		c = Coord{y: s.y - yDif, x: s.x - xDif}

		if inbound(c, width, height) {
			aNs[c] = true
		}
	}

	if len(crds) > 2 {
		calcPerms(crds[1:], aNs, height, width)
	}
}

func inbound(c Coord, width, height int) bool {
	return !(c.y < 0 || c.y >= height || c.x < 0 || c.x >= width)
}

func printRes(lines []string, pos map[Coord]bool) {

	for v := range pos {
		lines[v.y] = lines[v.y][:v.x] + "#" + lines[v.y][v.x+1:]
	}

	fmt.Println(strings.Join(lines, "\n"))

}
