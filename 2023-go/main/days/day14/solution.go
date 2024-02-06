package day14

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct {
	times int
}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := grids.NewGrid(lines)
	tilt(grid, grids.North[int]())
	return strconv.Itoa(score(grid))
}

func (s Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := grids.NewGrid(lines)
	spinCycle(grid, s.times)
	return strconv.Itoa(score(grid))
}

func spinCycle(grid grids.Grid[rune], times int) {
	if times == 0 {
		times = 1_000_000_000
	}

	cache := map[string]int{}

	cyclesCompleted := 0
	cyclingSince := 0
	for i := 1; i <= times; i++ {
		tilt(grid, grids.North[int]())
		tilt(grid, grids.West[int]())
		tilt(grid, grids.South[int]())
		tilt(grid, grids.East[int]())
		if when, seen := cache[grid.Dump()]; seen {
			cyclesCompleted = i
			cyclingSince = when
			break
		}
		cache[grid.Dump()] = i
	}

	cycleLength := cyclesCompleted - cyclingSince
	cyclesToEmulate := times - cyclesCompleted
	cycleOffsetAtEnd := cyclesToEmulate % cycleLength
	for i := 0; i < cycleOffsetAtEnd; i++ {
		tilt(grid, grids.North[int]())
		tilt(grid, grids.West[int]())
		tilt(grid, grids.South[int]())
		tilt(grid, grids.East[int]())
	}
}

func tilt(grid grids.Grid[rune], direction grids.Vector2D[int]) {
	switch direction {
	case grids.North[int]():
		for y := 0; y < grid.Height(); y++ {
			for x := 0; x < grid.Width(); x++ {
				pos := grids.Vector2D[int]{X: x, Y: y}
				if grid.AtPos(pos) == 'O' {
					moveRock(grid, pos, direction)
				}
			}
		}
	case grids.South[int](): // reverse Y
		for y := grid.Height() - 1; y >= 0; y-- {
			for x := 0; x < grid.Width(); x++ {
				pos := grids.Vector2D[int]{X: x, Y: y}
				if grid.AtPos(pos) == 'O' {
					moveRock(grid, pos, direction)
				}
			}
		}
	case grids.West[int]():
		for x := 0; x < grid.Width(); x++ {
			for y := 0; y < grid.Height(); y++ {
				pos := grids.Vector2D[int]{X: x, Y: y}
				if grid.AtPos(pos) == 'O' {
					moveRock(grid, pos, direction)
				}
			}
		}
	case grids.East[int]():
		// reverse X
		for x := grid.Width() - 1; x >= 0; x-- {
			for y := 0; y < grid.Height(); y++ {
				pos := grids.Vector2D[int]{X: x, Y: y}
				if grid.AtPos(pos) == 'O' {
					moveRock(grid, pos, direction)
				}
			}
		}
	}
}

func moveRock(grid grids.Grid[rune], position grids.Vector2D[int], direction grids.Vector2D[int]) {
	valid := func(pos grids.Vector2D[int]) bool {
		return pos.InBounds(0, 0, grid.Width()-1, grid.Height()-1) &&
			grid.AtPos(pos) != '#' &&
			grid.AtPos(pos) != 'O'
	}

	newPosition := position.Add(direction)
	for valid(newPosition) {
		grid.SetAtPos(position, '.')
		grid.SetAtPos(newPosition, 'O')
		position = newPosition
		newPosition = position.Add(direction)
	}
}

func score(grid grids.Grid[rune]) int {
	return util.Accumulate(util.TransformWithIndex(grid, func(row []rune, i int) int {
		multiplier := grid.Height() - i
		numRocks := strings.Count(string(row), "O")
		return multiplier * numRocks
	}), func(total, next int) int { return total + next })
}
