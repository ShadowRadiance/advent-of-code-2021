package day10

type Rover struct {
	pos    Position
	facing Direction
}

func (rover *Rover) move(n int) {
	rover.pos = rover.pos.Add(Position(rover.facing).ScalarProduct(n))
}

func (rover *Rover) arrow() rune {
	switch rover.facing {
	case North:
		return '↑'
	case South:
		return '↓'
	case East:
		return '→'
	case West:
		return '←'
	default:
		return '•'
	}
}
