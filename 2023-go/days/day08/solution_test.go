package day08

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				LLR
				
				AAA = (BBB, BBB)
				BBB = (AAA, ZZZ)
				ZZZ = (ZZZ, ZZZ)
			`),
			Expected: "6",
		},
		{
			Input: heredoc.Doc(`
				RL
				
				AAA = (BBB, CCC)
				BBB = (DDD, EEE)
				CCC = (ZZZ, GGG)
				DDD = (DDD, DDD)
				EEE = (EEE, EEE)
				GGG = (GGG, GGG)
				ZZZ = (ZZZ, ZZZ)
			`),
			Expected: "2",
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
				LR

				11A = (11B, XXX)
				11B = (XXX, 11Z)
				11Z = (11B, XXX)
				22A = (22B, XXX)
				22B = (22C, 22C)
				22C = (22Z, 22Z)
				22Z = (22B, 22B)
				XXX = (XXX, XXX)
			`),
			Expected: "6",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
