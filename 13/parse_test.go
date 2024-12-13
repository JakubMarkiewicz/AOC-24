package main

import (
	"reflect"
	"testing"
)

func TestGetInput(t *testing.T) {
	t.Run("sample input", func(t *testing.T) {
		input := "Button A: X+94, Y+34\nButton B: X+22, Y+67\nPrize: X=8400, Y=5400\n\nButton A: X+26, Y+66\nButton B: X+67, Y+21\nPrize: X=12748, Y=12176\n\nButton A: X+17, Y+86\nButton B: X+84, Y+37\nPrize: X=7870, Y=6450\n\nButton A: X+69, Y+23\nButton B: X+27, Y+71\nPrize: X=18641, Y=10279"
		got := getInput(input)
		want :=
			[]Game{
				{A: []int{94, 34}, B: []int{22, 67}, P: []int{8400, 5400}},
				{A: []int{26, 66}, B: []int{67, 21}, P: []int{12748, 12176}},
				{A: []int{17, 86}, B: []int{84, 37}, P: []int{7870, 6450}},
				{A: []int{69, 23}, B: []int{27, 71}, P: []int{18641, 10279}},
			}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)

		}
	})
}
