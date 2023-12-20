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
			Expected: "64",
		},
	}
	// for _, example := range examples {
	// 	actual := Solution{times: 1}.Part02(example.Input)
	// 	if strconv.Itoa(0+9+16+14+18+10+4+12+2+2) != actual {
	// 		t.Errorf("Expected: %v, got: %v", example.Expected, actual)
	// 	}
	// }
	// for _, example := range examples {
	// 	actual := Solution{times: 2}.Part02(example.Input)
	// 	if strconv.Itoa(0+9+0+7+18+10+8+9+4+4) != actual {
	// 		t.Errorf("Expected: %v, got: %v", example.Expected, actual)
	// 	}
	// }
	// for _, example := range examples {
	// 	actual := Solution{times: 3}.Part02(example.Input)
	// 	if strconv.Itoa(0+9+0+7+18+10+8+9+4+4) != actual {
	// 		t.Errorf("Expected: %v, got: %v", example.Expected, actual)
	// 	}
	// }
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
