package day11

import (
	"slices"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct {
	testing bool
}

func (Solution) Part01(input string) string {
	return solve(input, 2)
}

func (s Solution) Part02(input string) string {
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

	if s.testing {
		return solve(input, 10)
	} else {
		return solve(input, 1_000_000)
	}

}

func solve(input string, expansionMultiplier int) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	galaxies := readGalaxies(lines)
	minX, maxX, minY, maxY := determineBounds(galaxies)
	rowsToExpand, colsToExpand := determineEmptySpace(galaxies, minX, maxX, minY, maxY)
	galaxies = expandUniverse(galaxies, rowsToExpand, colsToExpand, expansionMultiplier-1)
	pairLengths := determinePairLengths(galaxies)
	sum := util.Accumulate(pairLengths, func(a int, b int) int { return a + b })
	return strconv.Itoa(sum)
}

func readGalaxies(lines []string) []grids.Position {
	galaxies := make([]grids.Position, 0)
	for row := 0; row < len(lines); row++ {
		if len(lines[row]) == 0 {
			continue
		}
		for col := 0; col < len(lines[0]); col++ {
			if lines[row][col] == '#' {
				galaxies = append(galaxies, grids.Position{X: col, Y: row})
			}
		}
	}
	return galaxies
}

func determineBounds(galaxies []grids.Position) (minX, maxX, minY, maxY int) {
	byX := func(a, b grids.Position) int { return a.X - b.X }
	byY := func(a, b grids.Position) int { return a.Y - b.Y }

	minX = slices.MinFunc(galaxies, byX).X
	maxX = slices.MaxFunc(galaxies, byX).X
	minY = slices.MinFunc(galaxies, byY).Y
	maxY = slices.MaxFunc(galaxies, byY).Y

	return
}

func determineEmptySpace(galaxies []grids.Position, minX, maxX, minY, maxY int) (rowsToExpand, colsToExpand []int) {
	for row := minY; row <= maxY; row++ {
		if util.None(galaxies, func(galaxy grids.Position) bool { return galaxy.Y == row }) {
			rowsToExpand = append(rowsToExpand, row)
		}
	}
	for col := minX; col <= maxX; col++ {
		if util.None(galaxies, func(galaxy grids.Position) bool { return galaxy.X == col }) {
			colsToExpand = append(colsToExpand, col)
		}
	}
	return
}

func expandUniverse(galaxies []grids.Position, rows, cols []int, amount int) []grids.Position {
	// offset positions from the bottom to not mess up rows indexing
	if last := len(rows) - 1; last >= 0 {
		for i, rowIndex := last, rows[0]; i >= 0; i-- {
			rowIndex = rows[i]

			for galIndex, galaxy := range galaxies {
				if galaxy.Y >= rowIndex {
					galaxies[galIndex].Y += amount
				}
			}
		}
	}

	// insert cols from the right to not mess up indexing
	if last := len(cols) - 1; last >= 0 {
		for i, colIndex := last, cols[0]; i >= 0; i-- {
			colIndex = cols[i]

			for galIndex, galaxy := range galaxies {
				if galaxy.X >= colIndex {
					galaxies[galIndex].X += amount
				}
			}
		}
	}
	return galaxies
}

func determinePairLengths(galaxies []grids.Position) []int {
	pairLengths := make([]int, 0)
	for i := 0; i < len(galaxies); i++ {
		galaxyA := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			galaxyB := galaxies[j]
			pairLengths = append(pairLengths, galaxyA.ManhattanDistance(galaxyB))
		}
	}
	return pairLengths
}
