package main

import (
	"strings"
)

func B(data []Robot, height, width int) {

	i := 0
	for {
		i++
		for dI, r := range data {
			data[dI] = moveRobotA(r, 1, height, width)

		}

		p := printB(data, height, width)
		if strings.Contains(p, "#############") {
			println(i)
			println(p)
			break
		}
	}

}

func printB(rs []Robot, h, w int) string {
	res := [][]string{}

	for i := 0; i < h; i++ {
		res = append(res, strings.Split(strings.Repeat(".", w), ""))
	}

	for _, r := range rs {
		res[r.pos.y][r.pos.x] = "#"
	}

	var resStr strings.Builder

	for _, r := range res {
		resStr.WriteString(strings.Join(r, ""))
		resStr.WriteString("\n")
	}

	return resStr.String()
}
