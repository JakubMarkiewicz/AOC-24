package main

import (
	"fmt"
	"strconv"
	"strings"
)

var cache = map[string]int{}

func B(stones []int) {
	res := 0

	for _, stone := range stones {
		res += blinkBT(stone, 75)
	}

	fmt.Println(res)
}

func blinkBT(stone, count int) int {
	key := buildKey(stone, count)
	if v, prs := cache[key]; prs {
		return v
	}

	if count == 0 {
		return 1
	}

	c := 0

	if stone == 0 {
		c += blinkBT(1, count-1)

	} else if str := strconv.Itoa(stone); len(str)%2 == 0 {
		left, _ := strconv.Atoi(str[:len(str)/2])
		right, _ := strconv.Atoi(str[len(str)/2:])
		c += blinkBT(left, count-1)
		c += blinkBT(right, count-1)

	} else {
		c += blinkBT(stone*2024, count-1)
	}

	cache[key] = c

	return c
}

var keyBuilder strings.Builder

func buildKey(stone, count int) string {
	keyBuilder.Reset()
	keyBuilder.WriteString(strconv.Itoa(stone))
	keyBuilder.WriteRune('-')
	keyBuilder.WriteString(strconv.Itoa(count))
	return keyBuilder.String()
}
