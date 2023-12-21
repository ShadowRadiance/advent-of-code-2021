package day13

import (
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
	reflections := util.Transform(patterns, func(pattern grids.Grid[rune]) *Reflection {
		return findReflections(pattern)[0]
	})
	reflectionScores := util.Transform(reflections, func(reflection *Reflection) int {
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
	// Every mirror has exactly one smudge:
	// exactly one . or # should be the opposite type.
	//
	// In each pattern, locate and fix the smudge that
	// causes a different reflection line to be valid.

	patterns := parsePatterns(lines)
	reflections := util.Transform(patterns, func(pattern grids.Grid[rune]) *Reflection {
		return findUnsmudgedReflection(pattern)
	})
	reflectionScores := util.Transform(reflections, func(reflection *Reflection) int {
		if reflection.vertical {
			return reflection.linesBeforeReflection
		}
		return 100 * reflection.linesBeforeReflection
	})
	sum := util.Accumulate(reflectionScores, func(total, next int) int { return total + next })

	return strconv.Itoa(sum)
}

type Reflection struct {
	vertical              bool
	linesBeforeReflection int
}

func parsePatterns(lines []string) []grids.Grid[rune] {
	patterns := make([]grids.Grid[rune], 0)

	var pattern grids.Grid[rune]
	for len(lines) > 0 {
		pattern, lines = parsePattern(lines)
		patterns = append(patterns, pattern)
	}
	return patterns
}

func parsePattern(lines []string) (grids.Grid[rune], []string) {
	firstEmptyIndex := slices.Index(lines, "")
	if firstEmptyIndex == -1 {
		return grids.NewGrid(lines), []string{}
	}
	return grids.NewGrid(lines[:firstEmptyIndex]), lines[firstEmptyIndex+1:]
}

func findReflections(pattern grids.Grid[rune]) []*Reflection {
	reflections := make([]*Reflection, 0)

	for _, reflection := range findReflectionVerticals(pattern) {
		reflections = append(reflections, reflection)
	}
	for _, reflection := range findReflectionHorizontals(pattern) {
		reflections = append(reflections, reflection)
	}

	return reflections
}

func findUnsmudgedReflection(pattern grids.Grid[rune]) *Reflection {
	smudgyReflection := findReflections(pattern)[0]

	for y := 0; y < pattern.Height(); y++ {
		for x := 0; x < pattern.Width(); x++ {
			newPattern := pattern.Clone()
			switch newPattern[y][x] {
			case '.':
				newPattern[y][x] = '#'
			case '#':
				newPattern[y][x] = '.'
			}

			unsmudgedReflections := findReflections(newPattern)
			for _, unsmudgedReflection := range unsmudgedReflections {
				if *unsmudgedReflection != *smudgyReflection {
					return unsmudgedReflection
				}
			}
		}
	}

	return nil
}

func findReflectionVerticals(pattern grids.Grid[rune]) []*Reflection {
	reflections := make([]*Reflection, 0)
	for i := 1; i < pattern.Width(); i++ {
		if string(pattern.ColAt(i-1)) == string(pattern.ColAt(i)) {
			if reflectedAroundVertical(pattern, i-1, i) {
				reflections = append(reflections, &Reflection{vertical: true, linesBeforeReflection: i})
			}
		}
	}
	return reflections
}

func findReflectionHorizontals(pattern grids.Grid[rune]) []*Reflection {
	reflections := make([]*Reflection, 0)
	for i := 1; i < pattern.Height(); i++ {
		if string(pattern.RowAt(i-1)) == string(pattern.RowAt(i)) {
			if reflectedAroundHorizontal(pattern, i-1, i) {
				reflections = append(reflections, &Reflection{vertical: false, linesBeforeReflection: i})
			}
		}
	}
	return reflections
}

func reflectedAroundVertical(pattern grids.Grid[rune], beforeIndex, afterIndex int) bool {
	maxI := min(beforeIndex, pattern.Width()-afterIndex-1)

	for i := 0; i <= maxI; i++ {
		if string(pattern.ColAt(beforeIndex-i)) != string(pattern.ColAt(afterIndex+i)) {
			return false
		}
	}
	return true
}

func reflectedAroundHorizontal(pattern grids.Grid[rune], beforeIndex, afterIndex int) bool {
	maxI := min(beforeIndex, pattern.Height()-afterIndex-1)

	for i := 0; i <= maxI; i++ {
		if string(pattern.RowAt(beforeIndex-i)) != string(pattern.RowAt(afterIndex+i)) {
			return false
		}
	}
	return true
}
