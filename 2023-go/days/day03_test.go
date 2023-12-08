package days

import (
	"github.com/MakeNowJust/heredoc"
	"testing"
)

func TestDay03_Part01(t *testing.T) {
	tests := []Test{
		{
			input: heredoc.Doc(`
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..
			`),
			expected: "4361",
		},
	}
	for _, test := range tests {
		actual := Day03{}.Part01(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}
func TestDay03_Part02(t *testing.T) {
	tests := []Test{
		{
			input: heredoc.Doc(`
                467..114..
                ...*......
                ..35..633.
                ......#...
                617*......
                .....+.58.
                ..592.....
                ......755.
                ...$.*....
                .664.598..
			`),
			expected: "467835",
		},
	}
	for _, test := range tests {
		actual := Day03{}.Part02(test.input)
		if test.expected != actual {
			t.Errorf("expected: %v, got: %v", test.expected, actual)
		}
	}
}
