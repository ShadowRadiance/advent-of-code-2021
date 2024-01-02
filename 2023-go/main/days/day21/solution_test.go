package day21

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
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
	input := heredoc.Doc(`
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
	`)

	expected := map[int]string{
		6:    "16",
		10:   "50",
		50:   "1594",
		100:  "6536",
		500:  "167004",
		1000: "668697",
		5000: "16733044",
	}

	for steps, result := range expected {
		actual := Solution{steps: steps}.Part02(input)
		if result != actual {
			t.Errorf("Expected: %v, got: %v", result, actual)
		}
	}

}
