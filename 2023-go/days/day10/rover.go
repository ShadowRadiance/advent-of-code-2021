package day10

import "github.com/shadowradiance/advent-of-code/2023-go/util/grids"

type Rover struct {
	pos    grids.Position
	facing grids.Direction
}

func (rover *Rover) move(n int) {
	rover.pos = rover.pos.Add(grids.Position(rover.facing).ScalarProduct(n))
}

func (rover *Rover) arrow() rune {
	switch rover.facing {
	case grids.North:
		return '↑'
	case grids.South:
		return '↓'
	case grids.East:
		return '→'
	case grids.West:
		return '←'
	default:
		return '•'
	}
}
