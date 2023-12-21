package grids

type Direction = Vector2D

var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	West  = Direction{-1, 0}
	East  = Direction{1, 0}
)

func (v Direction) Reverse() Direction {
	return v.ScalarProduct(-1)
}

func (v Direction) RotateLeft() Direction {
	if v.Y == 0 {
		// East/West => North/South
		// +1,0/-1,0 => 0,-1/0,+1
		return Direction{X: 0, Y: -v.X}
	} else {
		// North/South => West/East
		// 0,-1/0,+1 => -1,0/+1,0
		return Direction{X: v.Y, Y: 0}
	}
}

func (v Direction) RotateRight() Direction {
	if v.Y == 0 {
		// East/West => South/North
		// +1,0/-1,0 => 0,+1/0,-1
		return Direction{X: 0, Y: v.X}
	} else {
		// North/South => East/West
		// 0,-1/0,+1 => +1,0/-1,0
		return Direction{X: -v.Y, Y: 0}
	}
}
