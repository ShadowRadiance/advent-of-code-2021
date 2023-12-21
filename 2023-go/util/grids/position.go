package grids

import "math"

type Position = Vector2D

func (v Position) InBounds(x1, y1, x2, y2 int) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2
}

func (v Position) OutOfBounds(x1, y1, x2, y2 int) bool {
	return !v.InBounds(x1, y1, x2, y2)
}

func (v Position) ManhattanDistance(otherPos Position) int {
	return int(
		math.Abs(float64(v.X-otherPos.X)) +
			math.Abs(float64(v.Y-otherPos.Y)))
}
