package day10

import "github.com/shadowradiance/advent-of-code/2023-go/util/grids"

type Rover struct {
	pos    grids.Vector2D[int]
	facing grids.Vector2D[int]
}

func (rover *Rover) move(n int) {
	rover.pos = rover.pos.Add(rover.facing.ScalarProduct(n))
}

func (rover *Rover) arrow() rune {
	switch rover.facing {
	case grids.North[int]():
		return '↑'
	case grids.South[int]():
		return '↓'
	case grids.East[int]():
		return '→'
	case grids.West[int]():
		return '←'
	default:
		return '•'
	}
}
