package grids

type Vector2D struct {
	X, Y int
}

func (v Vector2D) Add(dir Vector2D) Vector2D {
	return Vector2D{X: v.X + dir.X, Y: v.Y + dir.Y}
}

func (v Vector2D) ScalarProduct(n int) Vector2D {
	return Vector2D{X: v.X * n, Y: v.Y * n}
}
