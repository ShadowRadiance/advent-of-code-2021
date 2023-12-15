package day10

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				-L|F7
				7S-7|
				L|7||
				-L-J|
				L|-JF
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				-L|-F7
				7S--7|
				L|7-||
				-L--J|
				L|--JF
			`),
			Expected: "5",
		},
		{
			Input: heredoc.Doc(`
				7-F7-
				.FJ|7
				SJLL7
				|F--J
				LJ.LJ
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
			Input: heredoc.Doc(`
				...........
				.S-------7.
				.|F-----7|.
				.||.....||.
				.||.....||.
				.|L-7.F-J|.
				.|..|.|..|.
				.L--J.L--J.
				...........
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				..........
				.S------7.
				.|F----7|.
				.||....||.
				.||....||.
				.|L-7F-J|.
				.|II||II|.
				.L--JL--J.
				..........
			`),
			Expected: "4",
		},
		{
			Input: heredoc.Doc(`
				.F----7F7F7F7F-7....
				.|F--7||||||||FJ....
				.||.FJ||||||||L7....
				FJL7L7LJLJ||LJ.L-7..
				L--J.L7...LJS7F-7L7.
				....F-J..F7FJ|L7L7L7
				....L7.F7||L7|.L7L7|
				.....|FJLJ|FJ|F7|.LJ
				....FJL-7.||.||||...
				....L---J.LJ.LJLJ...
			`),
			Expected: "8",
		},
		{
			Input: heredoc.Doc(`
				FF7FSF7F7F7F7F7F---7
				L|LJ||||||||||||F--J
				FL-7LJLJ||||||LJL-77
				F--JF--7||LJLJ7F7FJ-
				L---JF-JLJ.||-FJLJJ7
				|F|F-JF---7F7-L7L|7|
				|FFJF7L7F-JF7|JL---7
				7-L-JL7||F7|L7F-7F7|
				L.L7LFJ|||||FJL7||LJ
				L7JLJL-JLJLJL--JLJ.L
			`),
			Expected: "10",
		},
	}
	for i, example := range examples {
		if i == 100 {
			break
		}
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
