package main

import (
	"fmt"
	"strings"
	// "time"
)

func B(board [][]string, moves string, pos Coord) int {

	for _, move := range moves {
		// fmt.Println("\033[2J")
		nextPos := getNextPos(pos, move, 1)
		if isInbound(nextPos, board) && board[nextPos.y][nextPos.x] != "#" {
			if move == '<' || move == '>' {
				moveBHorizontal(board, move, &pos)
			} else {
				moveBVeritcal(board, move, &pos)
			}
		}
		// fmt.Println(string(move))
		// fmt.Println(printBoardB(board, pos))
		// time.Sleep(40 * time.Millisecond)
	}

	fmt.Println(printBoardB(board, pos))

	return calculateB(board)
}

func calculateB(board [][]string) int {

	res := 0

	for y, line := range board {
		for x, v := range line {
			if v == "[" {
				res += min((y*100)+x, (y*100)+(len(line)-1-1))
			}
		}
	}
	return res
}

func printBoardB(board [][]string, pos Coord) string {
	var str strings.Builder

	for i, line := range board {
		s := strings.Join(line, "")
		if i == pos.y {
			s = s[:pos.x] + Color + " " + Reset + s[pos.x+1:]
		}
		str.WriteString(s)
		str.WriteString("\n")
	}

	return str.String()
}

type CoordB struct {
	y, x1, x2 int
}

func moveBVeritcal(board [][]string, move rune, pos *Coord) {
	visited := []CoordB{}
	path := []CoordB{}

	nextPos := getNextPos(*pos, move, 1)
	if board[nextPos.y][nextPos.x] == "[" || board[nextPos.y][nextPos.x] == "]" {
		path = append(path, NewCoordB(nextPos, board))
	}

	for len(path) > 0 {
		first := path[0]
		path = path[1:]

		l := getNextPos(Coord{first.x1, first.y}, move, 1)
		r := getNextPos(Coord{first.x2, first.y}, move, 1)

		if !isInbound(l, board) || !isInbound(r, board) {
			return
		}
		lItem := board[l.y][l.x]
		rItem := board[r.y][r.x]

		if lItem == "#" || rItem == "#" {
			return
		}

		if lItem == "[" || lItem == "]" {
			path = append(path, NewCoordB(Coord{l.x, l.y}, board))
		}
		if rItem == "[" {
			path = append(path, NewCoordB(Coord{r.x, r.y}, board))
		}

		visited = append(visited, first)
	}

	if len(visited) > 0 {
		for i := len(visited) - 1; i >= 0; i-- {
			v := visited[i]
			if move == '^' {
				board[v.y-1][v.x1] = "["
				board[v.y-1][v.x2] = "]"
			} else {
				board[v.y+1][v.x1] = "["
				board[v.y+1][v.x2] = "]"
			}
			board[v.y][v.x1] = "."
			board[v.y][v.x2] = "."
		}
		board[visited[0].y][visited[0].x1] = "."
		board[visited[0].y][visited[0].x2] = "."
	}

	if move == '^' {
		pos.y -= 1
	} else {
		pos.y += 1
	}

}

func NewCoordB(c Coord, board [][]string) CoordB {
	if board[c.y][c.x] == "[" {
		return CoordB{c.y, c.x, c.x + 1}
	}
	return CoordB{c.y, c.x - 1, c.x}
}

func moveBHorizontal(board [][]string, move rune, pos *Coord) {
	path := []Coord{getNextPos(*pos, move, 1)}

	for {
		itemPos := path[len(path)-1]

		if !isInbound(itemPos, board) {
			return
		}

		item := board[itemPos.y][itemPos.x]

		if item == "." {
			// close + move
			for i := len(path) - 2; i >= 0; i-- {
				next := path[i+1]
				cur := path[i]
				board[next.y][next.x] = board[cur.y][cur.x]
			}
			board[path[0].y][path[0].x] = "."
			pos.y = path[0].y
			pos.x = path[0].x
			return
		} else if item == "#" {
			// can't move
			return
		} else {
			path = append(path, getNextPos(itemPos, move, 1), getNextPos(itemPos, move, 2))
		}
	}
}
