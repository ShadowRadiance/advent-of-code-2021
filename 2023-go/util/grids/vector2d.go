package grids

import "github.com/shadowradiance/advent-of-code/2023-go/util/constraints"

type Vector2D[T constraints.Signed] struct {
	X, Y T
}

func (v Vector2D[T]) Add(dir Vector2D[T]) Vector2D[T] {
	return Vector2D[T]{X: v.X + dir.X, Y: v.Y + dir.Y}
}

func (v Vector2D[T]) ScalarProduct(n T) Vector2D[T] {
	return Vector2D[T]{X: v.X * n, Y: v.Y * n}
}
