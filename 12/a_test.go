package main

import "testing"

func TestA(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		data := []string{
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		}
		got := A(data)
		want := 140

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("second example", func(t *testing.T) {
		data := []string{
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		}
		got := A(data)
		want := 772

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("third example", func(t *testing.T) {
		data := []string{
			"RRRRIICCFF",
			"RRRRIICCCF",
			"VVRRRCCFFF",
			"VVRCCCJFFF",
			"VVVVCJJCFE",
			"VVIVCCJJEE",
			"VVIIICJJEE",
			"MIIIIIJJEE",
			"MIIISIJEEE",
			"MMMISSJEEE",
		}
		got := A(data)
		want := 1930

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
