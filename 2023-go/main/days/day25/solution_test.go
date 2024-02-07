package day25

import (
	"testing"

	"github.com/MakeNowJust/heredoc"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

func TestSolution_Part01(t *testing.T) {
	examples := []util.TestExample{
		{
			Input: heredoc.Doc(`
				jqt: rhn xhk nvd
				rsh: frs pzl lsr
				xhk: hfx
				cmg: qnr nvd lhk bvb
				rhn: xhk bvb hfx
				bvb: xhk hfx
				pzl: lsr hfx nvd
				qnr: nvd
				ntq: jqt hfx bvb xhk
				nvd: lhk
				lsr: lhk
				rzs: qnr cmg lsr rsh
				frs: qnr lhk lsr`),
			Expected: "54",
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
				jqt: rhn xhk nvd
				rsh: frs pzl lsr
				xhk: hfx
				cmg: qnr nvd lhk bvb
				rhn: xhk bvb hfx
				bvb: xhk hfx
				pzl: lsr hfx nvd
				qnr: nvd
				ntq: jqt hfx bvb xhk
				nvd: lhk
				lsr: lhk
				rzs: qnr cmg lsr rsh
				frs: qnr lhk lsr`),
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
