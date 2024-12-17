package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

func main() {
	next := flag.Bool("next", false, "is second part")
	flag.Parse()

	if *next {
		board, moves, pos := parseInputB()
		fmt.Println(B(board, moves, pos))
	} else {
		board, moves, pos := parseInput()
		fmt.Println(A(board, moves, pos))
	}
}

type Coord struct {
	x, y int
}

func parseInput() ([][]string, string, Coord) {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	parsingMoves := false
	var moves string
	var board [][]string
	var pos Coord

	// skip first line
	sc.Scan()

	y := 1
	for sc.Scan() {
		line := sc.Text()

		x := strings.Index(line, "@")
		if x > 0 {
			pos.x = x - 1
			pos.y = y - 1
			line = line[0:x] + "." + line[x+1:]
		}
		y++

		if line == "" {
			parsingMoves = true
			board = board[:len(board)-1]
			continue
		}

		if parsingMoves {
			moves += line
		} else {
			board = append(board, strings.Split(line[1:len(line)-1], ""))
		}
	}

	return board, moves, pos
}

func parseInputB() ([][]string, string, Coord) {
	file, err := os.OpenFile("./input.txt", os.O_RDONLY, fs.ModePerm)

	if err != nil {
		log.Fatalf("Error opening file %v", err)
	}

	defer file.Close()

	sc := bufio.NewScanner(file)

	parsingMoves := false
	var moves string
	var board [][]string
	var pos Coord

	y := 1
	for sc.Scan() {
		line := sc.Text()

		if line == "" {
			parsingMoves = true
			continue
		}

		if parsingMoves {
			moves += line
		} else {
			var modified string
			for _, c := range line {
				switch c {
				case '#':
					modified += "##"
				case 'O':
					modified += "[]"
				case '.':
					modified += ".."
				case '@':
					pos.y = y - 1
					pos.x = len(modified)
					// It's supposed to be @. but i'm keeping the pos in seperate value
					modified += ".."
				}
			}
			board = append(board, strings.Split(modified, ""))
		}
		y++
	}

	return board, moves, pos
}
