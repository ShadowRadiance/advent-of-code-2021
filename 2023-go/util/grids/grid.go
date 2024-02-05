package grids

import (
	"fmt"
	"slices"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Grid[T any] [][]T

func NewGrid(lines []string) Grid[rune] {
	lines = util.Filter(lines, func(s string) bool { return len(s) > 0 })
	runeLines := make([][]rune, len(lines))

	for i, line := range util.Filter(lines, func(s string) bool { return len(s) > 0 }) {
		runeLines[i] = []rune(line)
	}
	return runeLines
}

func New[T any](height, width int) Grid[T] {
	rows := make([][]T, height)

	for i := range rows {
		rows[i] = make([]T, width)
	}

	return rows

}

func (g *Grid[T]) At(x, y int) T             { return (*g)[y][x] }
func (g *Grid[T]) AtPos(pos Position[int]) T { return g.At(pos.X, pos.Y) }
func (g *Grid[T]) Height() int               { return len(*g) }
func (g *Grid[T]) Width() int {
	if g.Height() == 0 {
		return 0
	}
	return len((*g)[0])
}

func (g *Grid[T]) RowAt(y int) []T { return (*g)[y] }
func (g *Grid[T]) ColAt(x int) []T {
	col := make([]T, g.Height())
	for y := range *g {
		col[y] = g.At(x, y)
	}
	return col
}

func (g *Grid[T]) Dump() string {
	s := ""
	for _, row := range *g {
		for _, item := range row {
			if maybeChar, ok := any(item).(rune); ok {
				s += fmt.Sprintf("%c ", maybeChar)
			} else {
				s += fmt.Sprintf("%v ", item)
			}

		}
		s += "\n"
	}
	return s
}

func (g *Grid[T]) Clone() Grid[T] {
	var newGrid Grid[T] = make([][]T, g.Height())
	for row := 0; row < newGrid.Height(); row++ {
		newGrid[row] = make([]T, g.Width())
		for col := 0; col < newGrid.Width(); col++ {
			newGrid.SetAt(col, row, g.At(col, row))
		}
	}
	return newGrid
}

func (g *Grid[T]) InsertRow(y int, newRow []T) {
	*g = slices.Insert(*g, y, newRow)
}
func (g *Grid[T]) InsertCol(x int, newCol []T) {
	// increase the size of all of the rows by 1
	// move all the elements of each row at or above index x to the right
	for rowIndex := range *g {
		(*g)[rowIndex] = slices.Insert(g.RowAt(rowIndex), x, newCol[rowIndex])
	}
}

func (g *Grid[T]) SetAt(x, y int, item T) {
	(*g)[y][x] = item
}
func (g *Grid[T]) SetAtPos(position Position[int], item T) {
	g.SetAt(position.X, position.Y, item)
}

func (g *Grid[T]) Clear(item T) {
	for y, items := range *g {
		for x := range items {
			g.SetAt(x, y, item)
		}
	}
}
