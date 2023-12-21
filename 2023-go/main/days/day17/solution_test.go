package day17

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				2413432311323
				3215453535623
				3255245654254
				3446585845452
				4546657867536
				1438598798454
				4457876987766
				3637877979653
				4654967986887
				4564679986453
				1224686865563
				2546548887735
				4322674655533
			`),
			Expected: "102",
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
				2413432311323
				3215453535623
				3255245654254
				3446585845452
				4546657867536
				1438598798454
				4457876987766
				3637877979653
				4654967986887
				4564679986453
				1224686865563
				2546548887735
				4322674655533
			`),
			Expected: "94",
		},
		{
			Input: heredoc.Doc(`
				111111111111
				999999999991
				999999999991
				999999999991
				999999999991
			`),
			Expected: "71",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
