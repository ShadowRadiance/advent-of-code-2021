package day21

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct {
	steps int
}

type Grid = grids.Grid[rune]
type Position = grids.Position[int]

func (s Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	if s.steps == 0 {
		s.steps = 64
	}

	grid := grids.NewGrid(lines)
	elf := findStart(grid)

	reachablePositions := findReachablePositions(grid, elf, s.steps)

	return strconv.Itoa(len(reachablePositions))
}

func (s Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	if s.steps == 0 {
		s.steps = 26501365
	}

	grid := grids.NewGrid(lines)
	elf := findStart(grid)

	reachablePositions := findReachablePositions(grid, elf, s.steps)

	return strconv.Itoa(len(reachablePositions))
}

func findStart(grid Grid) (pos Position) {
	for row, chars := range grid {
		for col, char := range chars {
			if char == 'S' {
				pos = Position{X: col, Y: row}
				grid.SetAtPos(pos, '.')
				return
			}
		}
	}
	panic("No start")
}

func findReachablePositions(grid Grid, pos Position, steps int) (reachable map[Position]bool) {
	reachable = map[Position]bool{}

	if steps == 0 {
		return
	}
	if steps == 1 {
		possiblePositions := []Position{
			pos.Add(grids.North[int]()),
			pos.Add(grids.East[int]()),
			pos.Add(grids.South[int]()),
			pos.Add(grids.West[int]()),
		}
		for _, position := range possiblePositions {
			if !position.InBounds(0, 0, grid.Width()-1, grid.Height()-1) {
				continue
			}
			if grid.AtPos(position) == '#' {
				continue
			}
			reachable[position] = true
		}
		return
	}

	lowerReachable := findReachablePositions(grid, pos, steps-1)
	for position := range lowerReachable {
		for p := range findReachablePositions(grid, position, 1) {
			reachable[p] = true
		}
	}
	return
}
