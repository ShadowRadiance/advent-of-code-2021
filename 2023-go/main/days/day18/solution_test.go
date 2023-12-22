package day18

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				R 6 (#70c710)
				D 5 (#0dc571)
				L 2 (#5713f0)
				D 2 (#d2c081)
				R 2 (#59c680)
				D 2 (#411b91)
				L 5 (#8ceee2)
				U 2 (#caa173)
				L 1 (#1b58a2)
				U 2 (#caa171)
				R 2 (#7807d2)
				U 3 (#a77fa3)
				L 2 (#015232)
				U 2 (#7a21e3)
			`),
			Expected: "62",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part01(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
func TestSolution_Part02(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				R 6 (#70c710)
				D 5 (#0dc571)
				L 2 (#5713f0)
				D 2 (#d2c081)
				R 2 (#59c680)
				D 2 (#411b91)
				L 5 (#8ceee2)
				U 2 (#caa173)
				L 1 (#1b58a2)
				U 2 (#caa171)
				R 2 (#7807d2)
				U 3 (#a77fa3)
				L 2 (#015232)
				U 2 (#7a21e3)
			`),
			Expected: "952408144115",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
