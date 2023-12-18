package day03

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
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

func (g Gear) ratio() (total int) {
	total = 1
	for _, p := range g.parts {
		total *= p.number
	}
	return
}

func (g Gear) contains(part *Part) bool {
	for _, p := range g.parts {
		if p == part {
			return true
		}
	}
	return false
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
				// numberStr runs from grids[r][c-len(numberStr)+1] to grids[r][c]
				potentialPart := Part{c - len(numberStr), c - 1, r, number}
				if _, found := grid.findSymbolAround(potentialPart); found {
					collectedParts = append(collectedParts, potentialPart)
				}
			}
		}
	}
	return
}

type Rectangle struct {
	lft, top, rgt, bot int
}

func (r Rectangle) expand(minX, minY, maxX, maxY int) Rectangle {
	return Rectangle{
		lft: util.MaxInt(r.lft-1, minX),
		top: util.MaxInt(r.top-1, minY),
		rgt: util.MinInt(r.rgt+1, maxX),
		bot: util.MinInt(r.bot+1, maxY),
	}
}

func (g Grid) findSymbolAround(part Part) (rune, bool) {
	// we need to look around that location for a symbol
	//
	// ???????????
	// ?numberStr?
	// ???????????
	//
	// note that we cannot look off the edges of the grids

	// push out the boundaries, respect the edges
	partRect := Rectangle{part.x1, part.y, part.x2, part.y}
	rect := partRect.expand(0, 0, g.numCols()-1, g.numRows()-1)

	for x := rect.lft; x <= rect.rgt; x++ {
		for y := rect.top; y <= rect.bot; y++ {
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

func buildPartsGrid(grid Grid, parts []Part) (partGrid [][]*Part) {
	// build an empty grids
	partGrid = make([][]*Part, grid.numRows())
	for i := 0; i < grid.numRows(); i++ {
		partGrid[i] = make([]*Part, grid.numCols())
	}
	// put references to all the parts on the grids
	for idx, part := range parts {
		for i := part.x1; i <= part.x2; i++ {
			partGrid[part.y][i] = &parts[idx]
		}
	}
	return
}

func makeGear(x, y int, partsGrid [][]*Part) (gear Gear) {
	gear = Gear{x, y, make([]*Part, 0)}

	rect := Rectangle{x, y, x, y}.expand(0, 0, len(partsGrid[0])-1, len(partsGrid)-1)

	for r := rect.top; r <= rect.bot; r++ {
		for c := rect.lft; c <= rect.rgt; c++ {
			if r != y || c != x {
				var part = partsGrid[r][c]
				if part != nil && !gear.contains(part) {
					gear.parts = append(gear.parts, part)
				}
			}
		}
	}

	return
}

func findGears(grid Grid, partsGrid [][]*Part) (gears []Gear) {
	gears = make([]Gear, 0)

	for r := 0; r < grid.numRows(); r++ {
		for c := 0; c < grid.numCols(); c++ {
			if grid.at(c, r) == '*' {
				gear := makeGear(c, r, partsGrid)
				if len(gear.parts) == 2 {
					gears = append(gears, gear)
				}
			}
		}
	}

	return
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
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}
	grid := buildGrid(lines)
	collectedParts := collectParts(grid)
	partsGrid := buildPartsGrid(grid, collectedParts)
	gears := findGears(grid, partsGrid)

	total := 0
	for _, gear := range gears {
		total += gear.ratio()
	}

	return strconv.Itoa(total)
}
