package main

import (
	"reflect"
	"testing"
)

func TestMoveRobotA(t *testing.T) {
	t.Run("sample input", func(t *testing.T) {
		data := Robot{Coord{2, 4}, Coord{2, -3}}
		got := moveRobotA(data, 5, 7, 11)
		want := Robot{Coord{1, 3}, Coord{2, -3}}

		if !reflect.DeepEqual(got, want) {
			t.Fatalf("got %v, want %v", got, want)
		}
	})
}

func TestJump(t *testing.T) {
	t.Run("upper bound", func(t *testing.T) {
		got := jump(5, 3)
		want := 2

		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
	t.Run("bottom bound", func(t *testing.T) {
		got := jump(-1, 3)
		want := 2
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
	t.Run("bottom bound overflow", func(t *testing.T) {
		got := jump(-11, 7)
		want := 3
		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
	t.Run("t", func(t *testing.T) {
		got := jump(-3, 2)
		want := 1

		if got != want {
			t.Fatalf("got %d, want %d", got, want)
		}
	})
}

func TestDetermineQuadrant(t *testing.T) {
	t.Run("sample valid", func(t *testing.T) {
		got, _ := determineQuadrant(Coord{0, 2}, 8, 11)
		want := 0

		if got != want {

			t.Fatalf("got %d, want %d", got, want)
		}
	})
	t.Run("sample invalid", func(t *testing.T) {
		got, err := determineQuadrant(Coord{0, 3}, 8, 11)

		if err == nil {
			t.Fatalf("expected error, got %d", got)
		}
	})
}
