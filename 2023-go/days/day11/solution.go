package day11

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := grids.NewGrid(lines)
	grid = expandGrid(grid)
	pairLengths := determinePairLengths(grid)
	sum := util.Accumulate(pairLengths, func(a int, b int) int { return a + b })

	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	// obviously we cannot ACTUALLY brute force this part
	// as we'd have to expand each row and column to insane amounts of space

	// if we move from an "Actual Grid" to a "Sparse Grid" (or simply a list
	// of galaxies -- Positions), we should be able to offset just those positions
	// rather than caring about all the empty space.

	// Step 1: apply the changes / refactor Part 1 to use this new approach
	// Step 2: verify it works with the test data in Part 1
	// Step 3: apply it to the test data in Part 2
	// Step 4: apply it to the real data in Part 2 and submit

	return "PENDING"
}

func expandGrid(grid grids.Grid) grids.Grid {
	rowIndicesNeedingExpansion := make([]int, 0)
	colIndicesNeedingExpansion := make([]int, 0)
	for y := 0; y < grid.Height(); y++ {
		if util.All(grid[y], func(value rune) bool { return value == '.' }) {
			rowIndicesNeedingExpansion = append(rowIndicesNeedingExpansion, y)
		}
	}
	for x := 0; x < grid.Width(); x++ {
		columnHasGalaxy := false
		for y := 0; y < grid.Height(); y++ {
			if grid.At(x, y) == '#' {
				columnHasGalaxy = true
				break
			}
		}
		if !columnHasGalaxy {
			colIndicesNeedingExpansion = append(colIndicesNeedingExpansion, x)
		}
	}

	newGrid := grid.Clone()

	// insert rows from the bottom to not mess up indexing
	if last := len(rowIndicesNeedingExpansion) - 1; last >= 0 {
		for i, rowIndex := last, rowIndicesNeedingExpansion[0]; i >= 0; i-- {
			rowIndex = rowIndicesNeedingExpansion[i]
			newGrid.InsertRow(rowIndex, emptySpaceRow(newGrid.Width()))
		}
	}

	// insert cols from the right to not mess up indexing
	if last := len(colIndicesNeedingExpansion) - 1; last >= 0 {
		for i, colIndex := last, colIndicesNeedingExpansion[0]; i >= 0; i-- {
			colIndex = colIndicesNeedingExpansion[i]
			newGrid.InsertCol(colIndex, emptySpaceCol(newGrid.Height()))
		}
	}
	return newGrid
}

func emptySpaceRow(width int) []rune {
	return []rune(strings.Repeat(".", width))
}

func emptySpaceCol(height int) []rune {
	return []rune(strings.Repeat(".", height))
}

func determinePairLengths(grid grids.Grid) []int {
	pairLengths := make([]int, 0)
	galaxies := findGalaxies(grid)
	for i := 0; i < len(galaxies); i++ {
		galaxyA := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			galaxyB := galaxies[j]
			pairLengths = append(pairLengths, galaxyA.ManhattanDistance(galaxyB))
		}
	}
	return pairLengths
}

func findGalaxies(grid grids.Grid) []grids.Position {
	var galaxies []grids.Position = make([]grids.Position, 0)
	for row, runeList := range grid {
		for col, runeChar := range runeList {
			if runeChar == '#' {
				galaxies = append(galaxies, grids.Position{X: col, Y: row})
			}
		}
	}
	return galaxies

}
