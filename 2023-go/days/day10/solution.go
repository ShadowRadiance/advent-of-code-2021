package day10

import (
	"slices"
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
			processInstruction(&rover, instruction)
			moves++
		}
	}
	displayGrid(grid, &rover)

	return strconv.Itoa(moves / 2)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := makeGrid(lines)
	start := findStart(grid)
	rover := Rover{start, allowableDirectionFrom(grid, start)}

	moves := 0

	var pipes Grid = make([][]rune, grid.height())
	for row := range pipes {
		pipes[row] = make([]rune, grid.width())
		for col := range pipes[row] {
			pipes[row][col] = '.'
		}
	}

	for {
		instruction := grid.atPos(rover.pos)
		pipes[rover.pos.y][rover.pos.x] = instruction
		if instruction == 'S' {
			if moves != 0 {
				break
			} else {
				rover.move(1)
				moves++
			}
		} else {
			processInstruction(&rover, instruction)
			moves++
		}
	}
	displayGrid(pipes, &rover)

	tripleGrid := expandGrid(pipes)
	floodFill(tripleGrid)

	inside := 0
	for row := range grid {
		for col := range grid[row] {
			if tripleGrid[row*3+2][col*3+2] == ' ' {
				tripleGrid[row*3+2][col*3+2] = 'I'
				inside++
			}
		}
	}
	displayGrid(tripleGrid, nil)

	return strconv.Itoa(inside)
}

func processInstruction(rover *Rover, instruction rune) {
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
		newPos = start.Add(Position(direction))
		if newPos.InBounds(0, 0, grid.width()-1, grid.height()-1) {
			instruction := grid.atPos(newPos)
			if connected(instruction, Direction(Position(direction).Reverse())) {
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

func displayGrid(grid Grid, rover *Rover) {
	grid2 := make([][]rune, grid.height())
	for row, runeLine := range grid {
		grid2[row] = make([]rune, grid.width())
		for col, runeChar := range runeLine {
			grid2[row][col] = runeChar
		}
	}

	if rover != nil {
		grid2[rover.pos.y][rover.pos.x] = rover.arrow()
	}

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

func expandGrid(grid Grid) Grid {
	newGrid := make([][]rune, grid.height()*3+2)
	for row := range newGrid {
		newGrid[row] = make([]rune, grid.width()*3+2)
		for col := range newGrid[row] {
			newGrid[row][col] = ' '
		}
	}

	for row := range grid {
		for col := range grid[row] {
			switch grid[row][col] {
			case '|':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune(" | ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune(" | ")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune(" | ")...)
			case '-':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune("   ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune("---")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune("   ")...)
			case 'L':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune(" L ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune(" LL")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune("   ")...)
			case 'J':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune(" J ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune("JJ ")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune("   ")...)
			case '7':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune("   ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune("77 ")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune(" 7 ")...)
			case 'F':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune("   ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune(" FF")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune(" F ")...)
			case '.':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune("   ")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune("   ")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune("   ")...)
			case 'S':
				slices.Replace(newGrid[row*3+1], col*3+1, col*3+4, []rune("SSS")...)
				slices.Replace(newGrid[row*3+2], col*3+1, col*3+4, []rune("SSS")...)
				slices.Replace(newGrid[row*3+3], col*3+1, col*3+4, []rune("SSS")...)
			default:
				panic("Unknown letter WTF")
			}
		}
	}
	return newGrid
}

func floodFill(grid Grid) Grid {
	start := Position{0, 0}
	sourceRune := grid.atPos(start) // should be ' '
	if sourceRune != 'O' {
		dfs(grid, start, sourceRune, 'O')
	}
	return grid
}

func dfs(grid Grid, pos Position, from, to rune) {
	if !pos.InBounds(0, 0, grid.width()-1, grid.height()-1) || grid.atPos(pos) != from {
		return
	}
	grid[pos.y][pos.x] = to
	dfs(grid, pos.Add(Position(West)), from, to)
	dfs(grid, pos.Add(Position(East)), from, to)
	dfs(grid, pos.Add(Position(North)), from, to)
	dfs(grid, pos.Add(Position(South)), from, to)
}
