package day24

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				19, 13, 30 @ -2,  1, -2
				18, 19, 22 @ -1, -1, -2
				20, 25, 34 @ -2, -2, -4
				12, 31, 28 @ -1, -2, -1
				20, 19, 15 @  1, -5, -3`),
			Expected: "2",
		},
	}
	for _, example := range examples {
		actual := Solution{testArea: &rect{
			x: 7,
			y: 7,
			h: 20,
			w: 20,
		}}.Part01(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
func TestSolution_Part02(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				19, 13, 30 @ -2,  1, -2
				18, 19, 22 @ -1, -1, -2
				20, 25, 34 @ -2, -2, -4
				12, 31, 28 @ -1, -2, -1
				20, 19, 15 @  1, -5, -3`),
			Expected: "47",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
