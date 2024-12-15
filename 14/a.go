package main

import (
	"fmt"
	"slices"
	"strconv"
)

var Yellow = fmt.Sprintf("\u001B[48;2;%d;%d;%dm", 255, 0, 0)

const Reset = "\033[0m"

func A(data []Robot, steps, height, width int) int {
	quadrants := make([]int, 4)

	dataN := []Robot{}

	for _, r := range data {

		r := moveRobotA(r, steps, height, width)
		dataN = append(dataN, r)
		quadrant, err := determineQuadrant(r.pos, height, width)
		if err != nil {
			continue
		}
		quadrants[quadrant] += 1
	}

	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

func print(rs []Robot, h, w int) string {
	res := [][]int{}

	for i := 0; i < h; i++ {
		res = append(res, slices.Repeat([]int{0}, w))
	}

	for _, r := range rs {
		res[r.pos.y][r.pos.x] += 1
	}

	resStr := ""

	for _, r := range res {
		for _, rS := range r {
			if rS > 0 {
				resStr += "#"
			} else {
				resStr += strconv.Itoa(rS)
			}
		}
		resStr += "\n"
	}

	return resStr
}

func determineQuadrant(c Coord, height, width int) (int, error) {
	horizontal := 0
	vertical := 0

	if c.x < width/2 {
		horizontal = 0
	} else if c.x >= width-width/2 {
		horizontal = 1
	} else {
		return 0, fmt.Errorf("Invalid")
	}

	if c.y < height/2 {
		vertical = 0
	} else if c.y >= height-height/2 {
		vertical = 1
	} else {
		return 0, fmt.Errorf("Invalid")
	}

	if horizontal == 0 && vertical == 0 {
		return 0, nil
	} else if horizontal == 1 && vertical == 0 {
		return 1, nil
	} else if horizontal == 0 && vertical == 1 {
		return 2, nil
	} else {
		return 3, nil
	}

}

func moveRobotA(r Robot, steps, h, w int) Robot {
	newR := Robot(r)

	newR.pos.x = jump(r.pos.x+(steps*r.move.x), w)
	newR.pos.y = jump(r.pos.y+(steps*r.move.y), h)

	return newR
}

func jump(v, max int) int {
	return (v%max + max) % max
}

func abs(v int) int {
	if v < 0 {
		return v * -1
	}

	return v
}
