package main

import (
	"fmt"
)

func B(lines []string, input map[rune][]Coord, height, width int) {
	antiNodes := make(map[Coord]bool)

	for _, crs := range input {
		calcPermsB(crs, antiNodes, height, width)
	}

	printRes(lines, antiNodes)
	fmt.Println(len(antiNodes))

}

func calcPermsB(crds []Coord, aNs map[Coord]bool, height, width int) {
	f := crds[0]
	aNs[f] = true

	for _, s := range crds[1:] {
		aNs[s] = true
		valid := true

		for i := 1; ; i++ {
			if !valid {
				break
			}
			yDif := (f.y - s.y) * i
			xDif := (f.x - s.x) * i

			c := Coord{y: f.y + yDif, x: f.x + xDif}

			if inbound(c, width, height) {
				aNs[c] = true
			} else {
				valid = false
			}

			c = Coord{y: s.y - yDif, x: s.x - xDif}

			if inbound(c, width, height) {
				aNs[c] = true
				valid = true
			}

		}

	}

	if len(crds) > 2 {
		calcPermsB(crds[1:], aNs, height, width)
	}
}
