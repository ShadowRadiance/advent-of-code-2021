package day10

import (
	"slices"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
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

	grid := grids.NewGrid(lines)
	start := findStart(grid)
	rover := Rover{start, allowableDirectionFrom(grid, start)}

	moves := 0
	for {
		instruction := grid.AtPos(rover.pos)
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

	grid := grids.NewGrid(lines)
	start := findStart(grid)
	rover := Rover{start, allowableDirectionFrom(grid, start)}

	moves := 0

	var pipes = make([][]rune, grid.Height())
	for row := range pipes {
		pipes[row] = make([]rune, grid.Width())
		for col := range pipes[row] {
			pipes[row][col] = '.'
		}
	}

	for {
		instruction := grid.AtPos(rover.pos)
		pipes[rover.pos.Y][rover.pos.X] = instruction
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
		if rover.facing == grids.South {
			rover.facing = grids.East
		} else if rover.facing == grids.West {
			rover.facing = grids.North
		}
	case 'J':
		if rover.facing == grids.South {
			rover.facing = grids.West
		} else if rover.facing == grids.East {
			rover.facing = grids.North
		}
	case '7':
		if rover.facing == grids.North {
			rover.facing = grids.West
		} else if rover.facing == grids.East {
			rover.facing = grids.South
		}
	case 'F':
		if rover.facing == grids.North {
			rover.facing = grids.East
		} else if rover.facing == grids.West {
			rover.facing = grids.South
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

func findStart(grid grids.Grid) grids.Position {
	for y := 0; y < grid.Height(); y++ {
		for x := 0; x < grid.Width(); x++ {
			if grid.At(x, y) == 'S' {
				return grids.Position{X: x, Y: y}
			}
		}
	}
	panic("Start not found!")
}

func allowableDirectionFrom(grid grids.Grid, start grids.Position) grids.Direction {
	var newPos grids.Position
	for _, direction := range []grids.Direction{grids.North, grids.East, grids.South, grids.West} {
		newPos = start.Add(grids.Position(direction))
		if newPos.InBounds(0, 0, grid.Width()-1, grid.Height()-1) {
			instruction := grid.AtPos(newPos)
			if connected(instruction, grids.Direction(grids.Position(direction).Reverse())) {
				return direction
			}
		}
	}
	panic("No allowable directions from position!")
}

func connected(instruction rune, direction grids.Direction) bool {
	switch instruction {
	case '|':
		return direction == grids.North || direction == grids.South
	case '-':
		return direction == grids.West || direction == grids.East
	case 'L':
		return direction == grids.North || direction == grids.East
	case 'J':
		return direction == grids.North || direction == grids.West
	case '7':
		return direction == grids.South || direction == grids.West
	case 'F':
		return direction == grids.South || direction == grids.East
	case '.':
		return false
	case 'S':
		panic("cannot read connected from S")
	default:
		panic("Unknown letter WTF")
	}
}

func displayGrid(grid grids.Grid, rover *Rover) {
	grid2 := make([][]rune, grid.Height())
	for row, runeLine := range grid {
		grid2[row] = make([]rune, grid.Width())
		for col, runeChar := range runeLine {
			grid2[row][col] = runeChar
		}
	}

	if rover != nil {
		grid2[rover.pos.Y][rover.pos.X] = rover.arrow()
	}

	for _, s := range grid2 {
		println(string(s))
	}
	println(strings.Repeat("-", grid.Width()))
}

func expandGrid(grid grids.Grid) grids.Grid {
	newGrid := make([][]rune, grid.Height()*3+2)
	for row := range newGrid {
		newGrid[row] = make([]rune, grid.Width()*3+2)
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

func floodFill(grid grids.Grid) grids.Grid {
	start := grids.Position{}
	sourceRune := grid.AtPos(start) // should be ' '
	if sourceRune != 'O' {
		dfs(grid, start, sourceRune, 'O')
	}
	return grid
}

func dfs(grid grids.Grid, pos grids.Position, from, to rune) {
	if !pos.InBounds(0, 0, grid.Width()-1, grid.Height()-1) || grid.AtPos(pos) != from {
		return
	}
	grid[pos.Y][pos.X] = to
	dfs(grid, pos.Add(grids.Position(grids.West)), from, to)
	dfs(grid, pos.Add(grids.Position(grids.East)), from, to)
	dfs(grid, pos.Add(grids.Position(grids.North)), from, to)
	dfs(grid, pos.Add(grids.Position(grids.South)), from, to)
}
