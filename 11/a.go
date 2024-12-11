package main

import (
	"fmt"
	"strconv"
)

func A(stones []int) {
	times := 25

	for times > 0 {
		stones = blinkA(stones)
		times--
	}

	fmt.Println(len(stones))

}

func blinkA(stones []int) []int {
	var afterBlink []int

	for _, s := range stones {
		if s == 0 {
			afterBlink = append(afterBlink, 1)
			continue
		}
		str := strconv.Itoa(s)
		if len(str)%2 == 0 {
			left, _ := strconv.Atoi(str[:len(str)/2])
			right, _ := strconv.Atoi(str[len(str)/2:])
			afterBlink = append(afterBlink, left)
			afterBlink = append(afterBlink, right)
		} else {
			afterBlink = append(afterBlink, s*2024)
		}
	}

	return afterBlink
}
