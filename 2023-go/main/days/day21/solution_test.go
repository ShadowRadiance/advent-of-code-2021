package day21

import (
	"fmt"
	"os"
	"path"
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

func readData(day int) string {
	// load input file based on day and part
	filename := fmt.Sprintf("day%02d.txt", day)
	if data, err := os.ReadFile(path.Join("..", "..", "data", filename)); err != nil {
		panic(err)
	} else {
		return string(data)
	}
}

func TestSolution_Part02(t *testing.T) {
	actual := Solution{}.Part02(readData(21))
	expected := "616951804315987"
	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}

}
