package day03

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"strconv"
	"strings"
)

type Solution struct{}

type Grid []string

func (g Grid) at(x, y int) rune {
	return rune(g[y][x])
}
func (g Grid) numRows() int { return len(g) }
func (g Grid) numCols() int { return len(g[0]) }

type Part struct {
	x1, x2, y int
	number    int
}

type Gear struct {
	x, y  int
	parts []*Part
}

func (g Gear) ratio(total int) {
	total = 1
	for _, p := range g.parts {
		total *= p.number
	}
	return
}

func buildGrid(lines []string) (grid Grid) {
	grid = make(Grid, 0, len(lines))
	for _, line := range lines {
		if len(line) != 0 {
			grid = append(grid, line)
		}
	}
	return
}

func collectParts(grid Grid) (collectedParts []Part) {
	collectedParts = make([]Part, 0)
	for r := 0; r < grid.numRows(); r++ {
		for c := 0; c < grid.numCols(); c++ {
			char := grid.at(c, r)
			numberStr := ""

			for strings.ContainsRune(util.NumberChars, char) {
				numberStr += string(char)
				c++
				if c < grid.numCols() {
					char = grid.at(c, r)
				} else {
					char = 0
				}
			}

			if len(numberStr) != 0 {
				number := util.ConvertNumeric(numberStr)
				// numberStr runs from grid[r][c-len(numberStr)+1] to grid[r][c]
				potentialPart := Part{c - len(numberStr), c - 1, r, number}
				if _, found := grid.findSymbolAround(potentialPart); found {
					collectedParts = append(collectedParts, potentialPart)
				}
			}
		}
	}
	return
}

func (g Grid) findSymbolAround(part Part) (rune, bool) {
	// we need to look around that location for a symbol
	//
	// ???????????
	// ?numberStr?
	// ???????????
	//
	// note that we cannot look off the edges of the grid

	// push out the boundaries, respect the edges
	lft := util.MaxInt(part.x1-1, 0)
	top := util.MaxInt(part.y-1, 0)
	rgt := util.MinInt(part.x2+1, g.numCols()-1)
	bot := util.MinInt(part.y+1, g.numRows()-1)

	for x := lft; x <= rgt; x++ {
		for y := top; y <= bot; y++ {
			if x < part.x1 || x > part.x2 || y != part.y { // don't check inside the number
				char := g.at(x, y)
				if !strings.ContainsRune("0123456789.", char) {
					return char, true
				}
			}
		}
	}
	return 0, false
}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}
	grid := buildGrid(lines)
	collectedParts := collectParts(grid)

	sumOfPartNumbers := 0
	for _, part := range collectedParts {
		sumOfPartNumbers += part.number
	}

	return strconv.Itoa(sumOfPartNumbers)
}

func (Solution) Part02(input string) string {
	if len(input) == 0 {
		return "NO DATA"
	}
	return "PENDING"
}
