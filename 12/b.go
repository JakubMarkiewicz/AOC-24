package main

import (
	"slices"
)

func B(lines []string) int {
	visited := make(map[string]bool)
	res := 0

	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[0]); x++ {
			crd := Coord{y, x}
			if _, prs := visited[crd.String()]; !prs {
				res += getPriceB(y, x, lines, visited)
			}
		}
	}

	return res
}

func getPriceB(startY, startX int, lines []string, visited map[string]bool) int {
	stack := []Coord{{startY, startX}}
	visited[Coord{startY, startX}.String()] = true

	fences := []map[int][]int{
		// top
		{},
		// bottom
		{},
		// left
		{},
		// right
		{},
	}

	area := 1
	val := lines[startY][startX]

	for len(stack) > 0 {
		popped := stack[0]
		stack = stack[1:]

		// check around
		aroundCrds := popped.Around()
		for i, v := range aroundCrds {
			_, prs := visited[v.String()]
			if !v.Inbound(lines) || lines[v.y][v.x] != val {
				if i < 2 {
					// handle top - bottom
					if _, prs := fences[i][v.y]; prs {
						fences[i][v.y] = append(fences[i][v.y], v.x)
						continue
					}
					fences[i][v.y] = []int{v.x}
				} else {
					// handle left - right
					if _, prs := fences[i][v.x]; prs {
						fences[i][v.x] = append(fences[i][v.x], v.y)
						continue
					}
					fences[i][v.x] = []int{v.y}
				}

				continue
			}
			sameVal := lines[v.y][v.x] == val
			if !prs && sameVal {
				visited[v.String()] = true
				stack = append(stack, v)
				area++
			}

		}
	}

	totalSides := 0

	for _, v := range fences {

		for _, vJ := range v {
			slices.Sort(vJ)
			if len(vJ) == 1 {
				totalSides++
				continue
			}
			for k := 1; k < len(vJ); k++ {
				if vJ[k]-vJ[k-1] != 1 {
					totalSides++
				}
				// handle last
				if k == len(vJ)-1 {
					totalSides++
				}
			}
		}

	}

	return area * totalSides

}
