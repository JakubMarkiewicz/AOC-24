package main

import (
	"fmt"
	"strings"
	"time"
)

var Reset = "\033[0m"
var Color = fmt.Sprintf("\u001B[48;2;%d;%d;%dm", 255, 0, 0)

func A(board [][]string, moves string, pos Coord) int {
	for _, move := range moves {
		nextPos := getNextPos(pos, move, 1)
		if isInbound(nextPos, board) && board[nextPos.y][nextPos.x] != "#" {
			moveA(board, move, &pos)
		}
		fmt.Println("\033[2J")
		fmt.Println(printBoardA(board, pos))
		fmt.Printf("%vGPS: %d%v\n", Color, calculate(board), Reset)
		time.Sleep(5 * time.Millisecond)

	}

	return calculate(board)
}

func calculate(board [][]string) int {
	res := 0
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == "O" {
				res += (100 * (y + 1)) + x + 1
			}
		}
	}
	return res
}

func printBoardA(board [][]string, pos Coord) string {
	var str strings.Builder

	for i, line := range board {
		s := strings.Join(line, "")
		if i == pos.y {
			s = s[:pos.x] + Color + "@" + Reset + s[pos.x+1:]
		}
		str.WriteString(s)
		str.WriteString("\n")
	}

	return str.String()
}

func moveA(board [][]string, move rune, pos *Coord) {
	path := []Coord{getNextPos(*pos, move, 1)}

	if board[path[0].y][path[0].x] == "." {
		pos.x = path[0].x
		pos.y = path[0].y
		return
	}

	for {
		lastPos := path[len(path)-1]
		nextPos := getNextPos(lastPos, move, 1)

		if !isInbound(nextPos, board) {
			return
		}
		nextItem := board[nextPos.y][nextPos.x]
		if nextItem == "#" {
			return
		} else if nextItem == "O" {
			path = append(path, nextPos)
		} else {
			board[nextPos.y][nextPos.x] = "O"
			board[path[0].y][path[0].x] = "."
			pos.x = path[0].x
			pos.y = path[0].y
			return
		}

	}

}

func getNextPos(pos Coord, move rune, count int) Coord {
	switch move {
	case '<':
		return Coord{pos.x - count, pos.y}
	case '>':
		return Coord{pos.x + count, pos.y}
	case '^':
		return Coord{pos.x, pos.y - count}
	default:
		return Coord{pos.x, pos.y + count}
	}
}

func isInbound(pos Coord, board [][]string) bool {
	if pos.x < 0 || pos.x >= len(board[0]) {
		return false
	}

	if pos.y < 0 || pos.y >= len(board) {
		return false
	}

	return true

}
