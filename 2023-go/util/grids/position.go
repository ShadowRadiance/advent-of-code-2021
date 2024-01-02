package grids

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/constraints"
)

type Position[T constraints.Signed] Vector2D[T]

func (v Position[T]) InBounds(x1, y1, x2, y2 T) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2
}

func (v Position[T]) OutOfBounds(x1, y1, x2, y2 T) bool {
	return !v.InBounds(x1, y1, x2, y2)
}

func (v Position[T]) ManhattanDistance(otherPos Position[T]) T {
	return T(util.Abs(v.X-otherPos.X) + util.Abs(v.Y-otherPos.Y))
}

func (v Position[T]) Add(dir Direction[T]) Position[T] {
	return Position[T](Vector2D[T](v).Add(Vector2D[T](dir)))
}

func (v Position[T]) Mod(width T, height T) Position[T] {
	newX, newY := v.X, v.Y
	for newX < 0 {
		newX = width + newX
	}
	for newY < 0 {
		newY = height + newY
	}
	for newX >= width {
		newX -= width
	}
	for newY >= height {
		newY -= height
	}

	return Position[T]{X: newX, Y: newY}
}
