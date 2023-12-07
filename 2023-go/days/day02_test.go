package days

import (
	"github.com/MakeNowJust/heredoc"
	"testing"
)

func TestDay02_Part01(t *testing.T) {
	tests := []Test{
		{
			input:    "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
			expected: "1",
		},
		{
			input:    "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
			expected: "2",
		},
		{
			input:    "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
			expected: "0", // because this game cannot be done with 12 red, 13 green, and 14 blue cubes
		},
		{
			input:    "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
			expected: "0", // because this game cannot be done with 12 red, 13 green, and 14 blue cubes
		},
		{
			input:    "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
			expected: "5",
		},
		{
			input: heredoc.Doc(`
				Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
				Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
				Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
				Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
				Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
			`),
			expected: "8",
		},
	}

	for _, test := range tests {
		actual := Day02{}.Part01(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}

func TestDay02_Part02(t *testing.T) {
	tests := []Test{}

	for _, test := range tests {
		actual := Day02{}.Part02(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}
