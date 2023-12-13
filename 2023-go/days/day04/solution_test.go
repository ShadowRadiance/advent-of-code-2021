package day04

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{Input: "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53", Expected: "8"},
		{Input: "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19", Expected: "2"},
		{Input: "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1", Expected: "2"},
		{Input: "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83", Expected: "1"},
		{Input: "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36", Expected: "0"},
		{Input: "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11", Expected: "0"},

		{
			Input: heredoc.Doc(`
        Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
			`),
			Expected: "13",
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
        Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
        Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
        Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
        Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
        Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
        Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
			`),
			Expected: "30",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
