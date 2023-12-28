package day05

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				seeds: 79 14 55 13
				
				seed-to-soil map:
				50 98 2
				52 50 48
				
				soil-to-fertilizer map:
				0 15 37
				37 52 2
				39 0 15
				
				fertilizer-to-water map:
				49 53 8
				0 11 42
				42 0 7
				57 7 4
				
				water-to-light map:
				88 18 7
				18 25 70
				
				light-to-temperature map:
				45 77 23
				81 45 19
				68 64 13
				
				temperature-to-humidity map:
				0 69 1
				1 0 69
				
				humidity-to-location map:
				60 56 37
				56 93 4
			`),
			Expected: "35",
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
				seeds: 79 14 55 13
				
				seed-to-soil map:
				50 98 2
				52 50 48
				
				soil-to-fertilizer map:
				0 15 37
				37 52 2
				39 0 15
				
				fertilizer-to-water map:
				49 53 8
				0 11 42
				42 0 7
				57 7 4
				
				water-to-light map:
				88 18 7
				18 25 70
				
				light-to-temperature map:
				45 77 23
				81 45 19
				68 64 13
				
				temperature-to-humidity map:
				0 69 1
				1 0 69
				
				humidity-to-location map:
				60 56 37
				56 93 4
			`),
			Expected: "46",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}

func TestSplitIntervals(t *testing.T) {
	seedInterval := Interval{
		Start: 82,
		Final: 92,
	}
	targetInterval := Interval{
		Start: 46,
		Final: 56,
	}
	translator := Map{
		name: "whatever",
		mappings: []Mapping{
			{60, 56, 37},
			{56, 93, 4},
		},
	}

	result := splitIntervals(seedInterval, targetInterval, translator)
	// { (82,91) => (46,55), (92,92) => (56, 56)

	if len(result) != 2 {
		t.Errorf("Expected 2 elements")
	}
	result1 := result[Interval{Start: 82, Final: 91}]
	expected1 := Interval{Start: 46, Final: 55}
	if result1 != expected1 {
		t.Errorf("Expected 1 elements to be (82,91) => (46,55)")
	}

	result2 := result[Interval{Start: 92, Final: 92}]
	expected2 := Interval{Start: 60, Final: 60}
	if result2 != expected2 {
		t.Errorf("Expected 1 elements to be (92,92) => (56, 56)")
	}
}
