package day12

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				???.### 1,1,3
			`),
			Expected: "1",
		},
		{
			Input: heredoc.Doc(`
				.??..??...?##. 1,1,3
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				?#?#?#?#?#?#?#? 1,3,1,6
			`),
			Expected: "1",
		},
		{
			Input: heredoc.Doc(`
				????.#...#... 4,1,1
			`),
			Expected: "1",
		},
		{
			Input: heredoc.Doc(`
				????.######..#####. 1,6,5
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				?###???????? 3,2,1
			`),
			Expected: "10",
		},
		{
			Input: heredoc.Doc(`
				???.### 1,1,3
				.??..??...?##. 1,1,3
				?#?#?#?#?#?#?#? 1,3,1,6
				????.#...#... 4,1,1
				????.######..#####. 1,6,5
				?###???????? 3,2,1
			`),
			Expected: "21",
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
				???.### 1,1,3
			`),
			Expected: "1",
		},
		{
			Input: heredoc.Doc(`
				.??..??...?##. 1,1,3
			`),
			Expected: "16384",
		},
		{
			Input: heredoc.Doc(`
				?#?#?#?#?#?#?#? 1,3,1,6
			`),
			Expected: "1",
		},
		{
			Input: heredoc.Doc(`
				????.#...#... 4,1,1
			`),
			Expected: "16",
		},
		{
			Input: heredoc.Doc(`
				????.######..#####. 1,6,5
			`),
			Expected: "2500",
		},
		{
			Input: heredoc.Doc(`
				?###???????? 3,2,1
			`),
			Expected: "506250",
		},
		{
			Input: heredoc.Doc(`
				???.### 1,1,3
				.??..??...?##. 1,1,3
				?#?#?#?#?#?#?#? 1,3,1,6
				????.#...#... 4,1,1
				????.######..#####. 1,6,5
				?###???????? 3,2,1
			`),
			Expected: "525152",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
