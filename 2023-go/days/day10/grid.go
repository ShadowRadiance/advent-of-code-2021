package day10

type Grid [][]rune

func (g Grid) at(x, y int) rune        { return g[y][x] }
func (g Grid) atPos(pos Position) rune { return g.at(pos.x, pos.y) }
func (g Grid) height() int             { return len(g) }
func (g Grid) width() int              { return len(g[0]) }
