package day22

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				1,0,1~1,2,1
				0,0,2~2,0,2
				0,2,3~2,2,3
				0,0,4~0,2,4
				2,0,5~2,2,5
				0,1,6~2,1,6
				1,1,8~1,1,9
			`),
			Expected: "5",
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
				1,0,1~1,2,1
				0,0,2~2,0,2
				0,2,3~2,2,3
				0,0,4~0,2,4
				2,0,5~2,2,5
				0,1,6~2,1,6
				1,1,8~1,1,9
			`),
			Expected: "7",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
