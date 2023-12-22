package grids

import "github.com/shadowradiance/advent-of-code/2023-go/util/constraints"

type Direction[T constraints.Signed] Vector2D[T]

func (v Direction[T]) ScalarProduct(distance T) Direction[T] {
	return Direction[T](Vector2D[T](v).ScalarProduct(distance))
}

func North[T constraints.Signed]() Direction[T] { return Direction[T]{0, -1} }
func South[T constraints.Signed]() Direction[T] { return Direction[T]{0, 1} }
func West[T constraints.Signed]() Direction[T]  { return Direction[T]{-1, 0} }
func East[T constraints.Signed]() Direction[T]  { return Direction[T]{1, 0} }

func (v Direction[T]) Reverse() Direction[T] {
	return Direction[T](Vector2D[T](v).ScalarProduct(-1))
}

func (v Direction[T]) RotateLeft() Direction[T] {
	if v.Y == 0 {
		// East/West => North/South
		// +1,0/-1,0 => 0,-1/0,+1
		return Direction[T]{X: 0, Y: -v.X}
	} else {
		// North/South => West/East
		// 0,-1/0,+1 => -1,0/+1,0
		return Direction[T]{X: v.Y, Y: 0}
	}
}

func (v Direction[T]) RotateRight() Direction[T] {
	if v.Y == 0 {
		// East/West => South/North
		// +1,0/-1,0 => 0,+1/0,-1
		return Direction[T]{X: 0, Y: v.X}
	} else {
		// North/South => East/West
		// 0,-1/0,+1 => +1,0/-1,0
		return Direction[T]{X: -v.Y, Y: 0}
	}
}
