package day02

import (
	"testing"

	"github.com/shadowradiance/advent-of-code/2023-go/util"

	"github.com/MakeNowJust/heredoc"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: "1",
		},
		{
			Input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: "2",
		},
		{
			Input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: "0", // because this game cannot be done with 12 red, 13 green, and 14 blue cubes
		},
		{
			Input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: "0", // because this game cannot be done with 12 red, 13 green, and 14 blue cubes
		},
		{
			Input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: "5",
		},
		{
			Input: heredoc.Doc(`
				Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
				Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
				Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
				Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
				Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`),
			Expected: "8",
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
			Input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			Expected: "48",
		},
		{
			Input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			Expected: "12",
		},
		{
			Input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			Expected: "1560",
		},
		{
			Input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			Expected: "630",
		},
		{
			Input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			Expected: "36",
		},
		{
			Input: heredoc.Doc(`
				Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
				Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
				Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
				Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
				Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`),
			Expected: "2286",
		},
	}

	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
