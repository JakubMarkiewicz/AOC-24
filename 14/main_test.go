package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	t.Run("single line", func(t *testing.T) {
		data := []string{
			"p=69,74 v=68,-98",
		}
		got := parseInput(data)
		want := []Robot{{pos: Coord{69, 74}, move: Coord{68, -98}}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
