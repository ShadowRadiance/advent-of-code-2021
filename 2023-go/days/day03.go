package days

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"strconv"
	"strings"
)

type Day03 struct{}

type Grid []string

func (g Grid) at(x, y int) rune {
	return rune(g[y][x])
}
func (g Grid) numRows() int { return len(g) }
func (g Grid) numCols() int { return len(g[0]) }

func (Day03) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	var grid Grid = make(Grid, 0, len(lines))
	for _, line := range lines {
		if len(line) != 0 {
			grid = append(grid, line)
		}
	}

	collectedPartNumbers := make([]string, 0)
	sumOfPartNumbers := 0
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
				// numberStr runs from grid[r][c-len(numberStr)+1] to grid[r][c]
				if _, found := grid.findSymbolAround(c-len(numberStr), r, c-1, r); found {
					collectedPartNumbers = append(collectedPartNumbers, numberStr)
					sumOfPartNumbers += util.ConvertNumeric(numberStr)
				}
			}
		}
	}

	//fmt.Println(collectedPartNumbers)

	return strconv.Itoa(sumOfPartNumbers)
}

func (g Grid) findSymbolAround(x1, y1, x2, y2 int) (rune, bool) {
	// we need to look around that location for a symbol
	//
	// ???????????
	// ?numberStr?
	// ???????????
	//
	// note that we cannot look off the edges of the grid

	// push out the boundaries, respect the edges
	x1 = util.MaxInt(x1-1, 0)
	y1 = util.MaxInt(y1-1, 0)
	x2 = util.MinInt(x2+1, g.numCols()-1)
	y2 = util.MinInt(y2+1, g.numRows()-1)

	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			char := g.at(x, y)
			if !strings.ContainsRune("0123456789.", char) {
				return char, true
			}
		}
	}
	return 0, false
}

func (Day03) Part02(input string) string {
	if len(input) == 0 {
		return "NO DATA"
	}
	return "PENDING"
}
