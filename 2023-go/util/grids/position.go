package grids

import "math"

type Position struct {
	X, Y int
}

func (pos Position) Add(dir Position) Position {
	return Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
}

func (pos Position) InBounds(x1, y1, x2, y2 int) bool {
	return pos.X >= x1 && pos.X <= x2 && pos.Y >= y1 && pos.Y <= y2
}

func (pos Position) OutOfBounds(x1, y1, x2, y2 int) bool {
	return !pos.InBounds(x1, y1, x2, y2)
}

func (pos Position) ScalarProduct(n int) Position {
	return Position{X: pos.X * n, Y: pos.Y * n}
}

func (pos Position) Reverse() Position {
	return pos.ScalarProduct(-1)
}

func (pos Position) ManhattanDistance(otherPos Position) int {
	return int(
		math.Abs(float64(pos.X-otherPos.X)) +
			math.Abs(float64(pos.Y-otherPos.Y)))
}
