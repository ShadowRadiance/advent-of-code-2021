package day20

import (
	"github.com/MakeNowJust/heredoc"
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"testing"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		// {
		// 	Input: heredoc.Doc(`
  //       broadcaster -> a, b, c
  //       %a -> b
  //       %b -> c
  //       %c -> inv
  //       &inv -> a
  //     `),
		// 	Expected: "32000000",
		// },
		{
			Input: heredoc.Doc(`
        broadcaster -> a
        %a -> inv, con
        &inv -> b
        %b -> con
        &con -> output
      `),
			Expected: "11687500",
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
			Input:    heredoc.Doc(``),
			Expected: "PENDING",
		},
	}
	for _, example := range examples {
		actual := Solution{}.Part02(example.Input)
		if example.Expected != actual {
			t.Errorf("Expected: %v, got: %v", example.Expected, actual)
		}
	}
}
