package day06

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				Time:      7
				Distance:  9
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				Time:      15
				Distance:  40
			`),
			Expected: "8",
		},
		{
			Input: heredoc.Doc(`
				Time:       30
				Distance:  200
			`),
			Expected: "9",
		},
		{
			Input: heredoc.Doc(`
				Time:      7  15   30
				Distance:  9  40  200
			`),
			Expected: "288",
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
				Time:      7  15   30
				Distance:  9  40  200
			`),
			Expected: "71503",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
