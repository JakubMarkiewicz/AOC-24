package main

import "testing"

func TestB(t *testing.T) {
	t.Run("first example", func(t *testing.T) {
		data := []string{
			"AAAA",
			"BBCD",
			"BBCC",
			"EEEC",
		}
		got := B(data)
		want := 80

		assertSame(t, got, want)
	})

	t.Run("second example", func(t *testing.T) {
		data := []string{
			"OOOOO",
			"OXOXO",
			"OOOOO",
			"OXOXO",
			"OOOOO",
		}
		got := B(data)
		want := 436

		assertSame(t, got, want)
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
		got := B(data)
		want := 1206

		assertSame(t, got, want)
	})

	t.Run("fourth example", func(t *testing.T) {
		data := []string{
			"EEEEE",
			"EXXXX",
			"EEEEE",
			"EXXXX",
			"EEEEE",
		}
		got := B(data)
		want := 236
		assertSame(t, got, want)
	})

	t.Run("fifth example", func(t *testing.T) {
		data := []string{
			"AAAAAA",
			"AAABBA",
			"AAABBA",
			"ABBAAA",
			"ABBAAA",
			"AAAAAA",
		}
		got := B(data)
		want := 368
		assertSame(t, got, want)
	})

}

func assertSame(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
