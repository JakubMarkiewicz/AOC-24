package main

import "testing"

func TestCountGameA(t *testing.T) {
	t.Run("single game", func(t *testing.T) {
		data := Game{
			[]int{94, 34}, []int{22, 67}, []int{8400, 5400},
		}

		got := countGameA(data)
		want := 280

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("single game - 2", func(t *testing.T) {
		data := Game{
			[]int{26, 66}, []int{67, 21}, []int{12748, 12176},
		}

		got := countGameA(data)
		want := 0

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("single game - 3", func(t *testing.T) {
		data := Game{
			[]int{17, 86}, []int{84, 37}, []int{7870, 6450},
		}

		got := countGameA(data)
		want := 200

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

}
