package grids

import "github.com/shadowradiance/advent-of-code/2023-go/util/constraints"

type Vector3D[T constraints.Signed] struct {
	X, Y, Z T
}

func (v Vector3D[T]) Add(dir Vector3D[T]) Vector3D[T] {
	return Vector3D[T]{X: v.X + dir.X, Y: v.Y + dir.Y, Z: v.Z + dir.Z}
}

func (v Vector3D[T]) ScalarProduct(n T) Vector3D[T] {
	return Vector3D[T]{X: v.X * n, Y: v.Y * n, Z: v.Z * n}
}
