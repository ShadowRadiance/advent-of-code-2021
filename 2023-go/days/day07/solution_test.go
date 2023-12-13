package day07

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input:    heredoc.Doc(`32T3K 765`),
			Expected: "765",
		},
		{
			Input:    heredoc.Doc(`T55J5 684`),
			Expected: "684",
		},
		{
			Input:    heredoc.Doc(`KK677 28`),
			Expected: "28",
		},
		{
			Input:    heredoc.Doc(`KTJJT 220`),
			Expected: "220",
		},
		{
			Input:    heredoc.Doc(`QQQJA 483`),
			Expected: "483",
		},
		{
			Input: heredoc.Doc(`
				32T3K 765
				T55J5 684
				KK677 28
				KTJJT 220
				QQQJA 483
			`),
			Expected: "6440",
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
				32T3K 765
				T55J5 684
				KK677 28
				KTJJT 220
				QQQJA 483
			`),
			Expected: "PENDING",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
