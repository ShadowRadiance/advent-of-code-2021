package day14

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				O....#....
				O.OO#....#
				.....##...
				OO.#O....O
				.O.....O#.
				O.#..O.#.#
				..O..#O..O
				.......O..
				#....###..
				#OO..#....
			`),
			Expected: "136",
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
				O....#....
				O.OO#....#
				.....##...
				OO.#O....O
				.O.....O#.
				O.#..O.#.#
				..O..#O..O
				.......O..
				#....###..
				#OO..#....
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
