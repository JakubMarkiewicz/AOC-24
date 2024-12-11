package main

import (
	"reflect"
	"testing"
)

func BenchmarkB(t *testing.B) {
	for i := 0; i < t.N; i++ {
		B([]int{125, 125})
	}
}

func TestBlinkBT(t *testing.T) {
	t.Run("1 iteration", func(t *testing.T) {
		got := blinkBT(125, 1)
		want := 1

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("2 iterations", func(t *testing.T) {
		got := blinkBT(125, 2)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("3 iterations", func(t *testing.T) {
		got := blinkBT(125, 3)
		want := 2

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
