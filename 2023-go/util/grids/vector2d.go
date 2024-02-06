package grids

import (
	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/constraints"
)

type Vector2D[T constraints.Signed] struct {
	X, Y T
}

func (v Vector2D[T]) Add(dir Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{X: v.X + dir.X, Y: v.Y + dir.Y}
}

func (v Vector2D[T]) ScalarProduct(n T) Vector2D[T] {
	return Vector2D[T]{X: v.X * n, Y: v.Y * n}
}

func North[T constraints.Signed]() Vector2D[T] { return Vector2D[T]{0, -1} }
func South[T constraints.Signed]() Vector2D[T] { return Vector2D[T]{0, 1} }
func West[T constraints.Signed]() Vector2D[T]  { return Vector2D[T]{-1, 0} }
func East[T constraints.Signed]() Vector2D[T]  { return Vector2D[T]{1, 0} }

func (v Vector2D[T]) Reverse() Vector2D[T] {
	return v.ScalarProduct(-1)
}

func (v Vector2D[T]) RotateLeft() Vector2D[T] {
	if v.Y == 0 {
		// East/West => North/South
		// +1,0/-1,0 => 0,-1/0,+1
		return Vector2D[T]{X: 0, Y: -v.X}
	} else {
		// North/South => West/East
		// 0,-1/0,+1 => -1,0/+1,0
		return Vector2D[T]{X: v.Y, Y: 0}
	}
}

func (v Vector2D[T]) RotateRight() Vector2D[T] {
	if v.Y == 0 {
		// East/West => South/North
		// +1,0/-1,0 => 0,+1/0,-1
		return Vector2D[T]{X: 0, Y: v.X}
	} else {
		// North/South => East/West
		// 0,-1/0,+1 => +1,0/-1,0
		return Vector2D[T]{X: -v.Y, Y: 0}
	}
}

func (v Vector2D[T]) InBounds(x1, y1, x2, y2 T) bool {
	return v.X >= x1 && v.X <= x2 && v.Y >= y1 && v.Y <= y2
}

func (v Vector2D[T]) OutOfBounds(x1, y1, x2, y2 T) bool {
	return !v.InBounds(x1, y1, x2, y2)
}

func (v Vector2D[T]) ManhattanDistance(otherPos Vector2D[T]) T {
	return T(util.Abs(v.X-otherPos.X) + util.Abs(v.Y-otherPos.Y))
}

func (v Vector2D[T]) ModWrap(width T, height T) Vector2D[T] {
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

	return Vector2D[T]{X: newX, Y: newY}
}
