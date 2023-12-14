package day09

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input:    `0 3 6 9 12 15`,
			Expected: "18",
		},
		{
			Input:    `1 3 6 10 15 21`,
			Expected: "28",
		},
		{
			Input:    `10 13 16 21 30 45`,
			Expected: "68",
		},
		{
			Input: heredoc.Doc(`
				0 3 6 9 12 15
				1 3 6 10 15 21
				10 13 16 21 30 45
			`),
			Expected: "114",
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
			Input:    `0 3 6 9 12 15`,
			Expected: "-3",
		},
		{
			Input:    `1 3 6 10 15 21`,
			Expected: "0",
		},
		{
			Input:    `10 13 16 21 30 45`,
			Expected: "5",
		},
		{
			Input: heredoc.Doc(`
				0 3 6 9 12 15
				1 3 6 10 15 21
				10 13 16 21 30 45
			`),
			Expected: "2",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
