package main

import (
	"strconv"
	"strings"
)

func A(lines []string) int {
	visited := make(map[string]bool)
	res := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			crd := Coord{y, x}
			if _, prs := visited[crd.String()]; !prs {
				res += getPriceA(y, x, lines, visited)
			}
		}
	}

	return res
}

func getPriceA(startY, startX int, lines []string, visited map[string]bool) int {
	stack := []Coord{{startY, startX}}
	visited[Coord{startY, startX}.String()] = true

	area := 1
	fences := 0

	val := lines[startY][startX]

	for len(stack) > 0 {
		popped := stack[0]
		stack = stack[1:]

		// check around
		aroundCrds := popped.Around()
		for _, v := range aroundCrds {
			_, prs := visited[v.String()]
			if !v.Inbound(lines) {
				fences++
				continue
			}
			sameVal := lines[v.y][v.x] == val
			if !prs && sameVal {
				visited[v.String()] = true
				stack = append(stack, v)
				area++
			}

			if !sameVal {
				fences++
			}
		}
	}

	return area * fences

}

type Coord struct {
	y, x int
}

func (c Coord) String() string {

	var str strings.Builder

	str.WriteString(strconv.Itoa(c.y))
	str.WriteRune('-')
	str.WriteString(strconv.Itoa(c.x))

	return str.String()
}

func (c Coord) Around() []Coord {
	return []Coord{
		{c.y - 1, c.x},
		{c.y + 1, c.x},
		{c.y, c.x - 1},
		{c.y, c.x + 1},
	}

}

func (c Coord) Inbound(src []string) bool {
	if c.y < 0 || c.y >= len(src) {
		return false
	}
	if c.x < 0 || c.x >= len(src[0]) {
		return false
	}

	return true
}
