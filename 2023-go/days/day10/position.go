package day10

type Position struct {
	x, y int
}

func (pos Position) Add(dir Position) Position {
	return Position{x: pos.x + dir.x, y: pos.y + dir.y}
}

func (pos Position) InBounds(x1, y1, x2, y2 int) bool {
	return pos.x >= x1 && pos.x <= x2 && pos.y >= y1 && pos.y <= y2
}

func (pos Position) OutOfBounds(x1, y1, x2, y2 int) bool {
	return !pos.InBounds(x1, y1, x2, y2)
}

func (pos Position) ScalarProduct(n int) Position {
	return Position{x: pos.x * n, y: pos.y * n}
}

func (pos Position) Reverse() Position {
	return pos.ScalarProduct(-1)
}
