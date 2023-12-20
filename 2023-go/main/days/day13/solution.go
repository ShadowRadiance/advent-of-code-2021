package day13

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	patterns := parsePatterns(lines)
	reflections := util.Transform(patterns, func(pattern grids.Grid) Reflection { return findReflection(pattern) })
	reflectionScores := util.Transform(reflections, func(reflection Reflection) int {
		if reflection.vertical {
			return reflection.linesBeforeReflection
		}
		return 100 * reflection.linesBeforeReflection
	})
	sum := util.Accumulate(reflectionScores, func(total, next int) int { return total + next })

	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

type Reflection struct {
	vertical              bool
	linesBeforeReflection int
}

func parsePatterns(lines []string) []grids.Grid {
	patterns := make([]grids.Grid, 0)

	var pattern grids.Grid
	for len(lines) > 0 {
		pattern, lines = parsePattern(lines)
		patterns = append(patterns, pattern)
	}
	return patterns
}

func parsePattern(lines []string) (grids.Grid, []string) {
	firstEmptyIndex := slices.Index(lines, "")
	if firstEmptyIndex == -1 {
		return grids.NewGrid(lines), []string{}
	}
	return grids.NewGrid(lines[:firstEmptyIndex]), lines[firstEmptyIndex+1:]
}

func findReflection(pattern grids.Grid) Reflection {
	if reflection, ok := findReflectionVertical(pattern); ok {
		return reflection
	}

	if reflection, ok := findReflectionHorizontal(pattern); ok {
		return reflection
	}

	panic(
		fmt.Sprintf("No reflections found in pattern %+v",
			util.Transform(pattern, func(row []rune) string {
				return string(row)
			})))
}

func findReflectionVertical(pattern grids.Grid) (Reflection, bool) {
	for i := 1; i < pattern.Width(); i++ {
		if string(pattern.ColAt(i-1)) == string(pattern.ColAt(i)) {
			if reflectedAroundVertical(pattern, i-1, i) {
				return Reflection{vertical: true, linesBeforeReflection: i}, true
			}
		}
	}
	return Reflection{}, false
}

func findReflectionHorizontal(pattern grids.Grid) (Reflection, bool) {
	for i := 1; i < pattern.Height(); i++ {
		if string(pattern.RowAt(i-1)) == string(pattern.RowAt(i)) {
			if reflectedAroundHorizontal(pattern, i-1, i) {
				return Reflection{vertical: false, linesBeforeReflection: i}, true
			}
		}
	}
	return Reflection{}, false
}

func reflectedAroundVertical(pattern grids.Grid, beforeIndex, afterIndex int) bool {
	maxI := min(beforeIndex, pattern.Width()-afterIndex-1)

	for i := 0; i <= maxI; i++ {
		if string(pattern.ColAt(beforeIndex-i)) != string(pattern.ColAt(afterIndex+i)) {
			return false
		}
	}
	return true
}

func reflectedAroundHorizontal(pattern grids.Grid, beforeIndex, afterIndex int) bool {
	maxI := min(beforeIndex, pattern.Height()-afterIndex-1)

	for i := 0; i <= maxI; i++ {
		if string(pattern.RowAt(beforeIndex-i)) != string(pattern.RowAt(afterIndex+i)) {
			return false
		}
	}
	return true
}
