package day21

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"testing"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				...........
				.....###.#.
				.###.##..#.
				..#.#...#..
				....#.#....
				.##..S####.
				.##..#...#.
				.......##..
				.##.#.####.
				.##..##.##.
				...........
			`),
			Expected: "16",
		},
	}
	for _, example := range examples {
		actual := Solution{steps: 6}.Part01(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
func TestSolution_Part02(t *testing.T) {
	examples := []util.TestExample{
		{
			Input:    heredoc.Doc(``),
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
