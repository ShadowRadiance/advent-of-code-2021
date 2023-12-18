package grids

import (
	"slices"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Grid [][]rune

func (g *Grid) At(x, y int) rune        { return (*g)[y][x] }
func (g *Grid) AtPos(pos Position) rune { return g.At(pos.X, pos.Y) }
func (g *Grid) Height() int             { return len(*g) }
func (g *Grid) Width() int              { return len((*g)[0]) }
func (g *Grid) RowAt(y int) []rune      { return (*g)[y] }

func NewGrid(lines []string) Grid {
	lines = util.Filter(lines, func(s string) bool { return len(s) > 0 })
	runeLines := make([][]rune, len(lines))

	for i, line := range util.Filter(lines, func(s string) bool { return len(s) > 0 }) {
		runeLines[i] = []rune(line)
	}
	return runeLines
}

func (g *Grid) Dump() {
	for _, s := range *g {
		println(string(s))
	}
}

func (g *Grid) Clone() Grid {
	var newGrid Grid = make([][]rune, g.Height())
	for row := 0; row < newGrid.Height(); row++ {
		newGrid[row] = make([]rune, g.Width())
		for col := 0; col < newGrid.Width(); col++ {
			newGrid[row][col] = g.At(col, row)
		}
	}
	return newGrid
}

func (g *Grid) InsertRow(y int, newRow []rune) {
	*g = slices.Insert(*g, y, newRow)
}

func (g *Grid) InsertCol(x int, newCol []rune) {
	// increase the size of all of the rows by 1
	// move all the elements of each row at or above index x to the right
	for rowIndex := range *g {
		(*g)[rowIndex] = slices.Insert(g.RowAt(rowIndex), x, newCol[rowIndex])
	}
}
