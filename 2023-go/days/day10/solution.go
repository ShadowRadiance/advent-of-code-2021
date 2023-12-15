package day10

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	// ••F7•  =>  •••••  =>  •••••  =>  •••••  =>  ••4••  =>  ••45•  =>  ••45•  =>  ••45•  =>  ••45•  =>
	// •FJ|•  =>  •••••  =>  •2•••  =>  •23••  =>  •23••  =>  •23••  =>  •236•  =>  •236•  =>  •236•  =>
	// SJ•L7  =>  01•••  =>  01•••  =>  01•••  =>  01•••  =>  01•••  =>  01•••  =>  01•78  =>  01•78  =>  8
	// |F--J  =>  1••••  =>  1••••  =>  1••••  =>  14•••  =>  145••  =>  1456•  =>  14567  =>  14567  =>
	// LJ•••  =>  •••••  =>  2••••  =>  23•••  =>  23•••  =>  23•••  =>  23•••  =>  23•••  =>  23•••  =>

	// •••••  =>  •••••  =>  •••••  =>  •••••  =>  •••••  =>
	// •S-7•  =>  •01••  =>  •012•  =>  •012•  =>  •012•  =>
	// •|•|•  =>  •1•••  =>  •1•••  =>  •1•3•  =>  •1•3•  =>  4
	// •L-J•  =>  •••••  =>  •2•••  =>  •23••  =>  •234•  =>
	// •••••  =>  •••••  =>  •••••  =>  •••••  =>  •••••  =>

	// the description makes it sound like pathfinding, but there is only one path
	// so lets follow it until we get back to the start and divide by two to get the "furthest steps away"

	grid := makeGrid(lines)
	start := findStart(grid)
	rover := Rover{start, allowableDirectionFrom(grid, start)}
	// displayGrid(grid, rover)

	moves := 0
	for {
		instruction := grid.atPos(rover.pos)
		if instruction == 'S' {
			if moves != 0 {
				break
			} else {
				rover.move(1)
				moves++
			}
		} else {
			rover.processInstruction(instruction)
			moves++
		}
	}
	// displayGrid(grid, rover)

	return strconv.Itoa(moves / 2)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

type Grid [][]rune

func (g Grid) at(x, y int) rune        { return g[y][x] }
func (g Grid) atPos(pos Position) rune { return g.at(pos.x, pos.y) }
func (g Grid) height() int             { return len(g) }
func (g Grid) width() int              { return len(g[0]) }

type Position struct {
	x, y int
}

type Direction Position

var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	West  = Direction{-1, 0}
	East  = Direction{1, 0}
)

type Rover struct {
	pos    Position
	facing Direction
}

func (pos Position) Add(dir Direction) Position {
	return Position{x: pos.x + dir.x, y: pos.y + dir.y}
}

func (pos Position) InBounds(x1, y1, x2, y2 int) bool {
	return pos.x >= x1 && pos.x <= x2 && pos.y >= y1 && pos.y <= y2
}

func (pos Position) OutOfBounds(x1, y1, x2, y2 int) bool {
	return !pos.InBounds(x1, y1, x2, y2)
}

func (dir Direction) ScalarProduct(n int) Direction {
	return Direction{x: dir.x * n, y: dir.y * n}
}

func (dir Direction) Reverse() Direction {
	return dir.ScalarProduct(-1)
}

func (rover *Rover) processInstruction(instruction rune) {
	switch instruction {
	case '|': // do nothing
	case '-': // do nothing
	case 'L':
		if rover.facing == South {
			rover.facing = East
		} else if rover.facing == West {
			rover.facing = North
		}
	case 'J':
		if rover.facing == South {
			rover.facing = West
		} else if rover.facing == East {
			rover.facing = North
		}
	case '7':
		if rover.facing == North {
			rover.facing = West
		} else if rover.facing == East {
			rover.facing = South
		}
	case 'F':
		if rover.facing == North {
			rover.facing = East
		} else if rover.facing == West {
			rover.facing = South
		}
	case '.':
		panic("Escaped the pipe somehow")
	case 'S':
		panic("Caller didn't resolve 'S' before passing")
	default:
		panic("Unknown letter WTF: " + string(instruction))
	}
	rover.move(1)
}

func (rover *Rover) move(n int) {
	rover.pos = rover.pos.Add(rover.facing.ScalarProduct(n))
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

func findStart(grid Grid) Position {
	for y := 0; y < grid.height(); y++ {
		for x := 0; x < grid.width(); x++ {
			if grid.at(x, y) == 'S' {
				return Position{x, y}
			}
		}
	}
	panic("Start not found!")
}

func allowableDirectionFrom(grid Grid, start Position) Direction {
	var newPos Position
	for _, direction := range []Direction{North, East, South, West} {
		newPos = start.Add(direction)
		if newPos.InBounds(0, 0, grid.width()-1, grid.height()-1) {
			instruction := grid.atPos(newPos)
			if connected(instruction, direction.Reverse()) {
				return direction
			}
		}
	}
	panic("No allowable directions from position!")
}

func connected(instruction rune, direction Direction) bool {
	switch instruction {
	case '|':
		return direction == North || direction == South
	case '-':
		return direction == West || direction == East
	case 'L':
		return direction == North || direction == East
	case 'J':
		return direction == North || direction == West
	case '7':
		return direction == South || direction == West
	case 'F':
		return direction == South || direction == East
	case '.':
		return false
	case 'S':
		panic("cannot read connected from S")
	default:
		panic("Unknown letter WTF")
	}
}

func displayGrid(grid Grid, rover Rover) {
	grid2 := make([][]rune, grid.height())
	for row, runeLine := range grid {
		grid2[row] = make([]rune, grid.width())
		for col, runeChar := range runeLine {
			grid2[row][col] = runeChar
		}
	}
	grid2[rover.pos.y][rover.pos.x] = rover.arrow()

	for _, s := range grid2 {
		println(string(s))
	}
	println(strings.Repeat("-", grid.width()))
}

func makeGrid(lines []string) Grid {
	lines = util.Filter(lines, func(s string) bool { return len(s) > 0 })
	runeLines := make([][]rune, len(lines))

	for i, line := range util.Filter(lines, func(s string) bool { return len(s) > 0 }) {
		runeLines[i] = []rune(line)
	}
	return runeLines
}
