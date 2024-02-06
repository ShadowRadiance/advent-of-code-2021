package day21

import (
	"strconv"
	"strings"

	aq "github.com/emirpasic/gods/queues/arrayqueue"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct {
	steps int
}

type Grid = grids.Grid[rune]
type Position = grids.Vector2D[int]

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

	numReachablePositions := fill(grid, elf, s.steps)
	return strconv.Itoa(numReachablePositions)
}

func (s Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	if s.steps == 0 {
		s.steps = 26_501_365
	}

	grid := grids.NewGrid(lines)
	elf := findStart(grid)

	util.Assert(grid.Width() == grid.Height(), "Grid is not square")
	size := grid.Width()
	util.Assert(size%2 == 1, "Grid width is not odd")
	util.Assert(elf.X == size/2 && elf.Y == size/2, "Didn't start in center")
	util.Assert(s.steps%size == size/2, "We can't reach the edge of the last grid")

	gridWidth := s.steps/size - 1
	numOddGrids := square(((gridWidth / 2) * 2) + 1)
	numEvenGrids := square(((gridWidth + 1) / 2) * 2)

	oddPoints := fill(grid, elf, size*2+1)
	evenPoints := fill(grid, elf, size*2)

	cornerNPoints := fill(grid, Position{X: elf.X, Y: size - 1}, size-1)
	cornerEPoints := fill(grid, Position{X: 0, Y: elf.Y}, size-1)
	cornerSPoints := fill(grid, Position{X: elf.X, Y: 0}, size-1)
	cornerWPoints := fill(grid, Position{X: size - 1, Y: elf.Y}, size-1)

	smallSteps := size/2 - 1
	smallNE := fill(grid, Position{X: 0, Y: size - 1}, smallSteps)
	smallSE := fill(grid, Position{X: 0, Y: 0}, smallSteps)
	smallNW := fill(grid, Position{X: size - 1, Y: 0}, smallSteps)
	smallSW := fill(grid, Position{X: size - 1, Y: size - 1}, smallSteps)

	largeSteps := 3*size/2 - 1
	largeNE := fill(grid, Position{X: 0, Y: size - 1}, largeSteps)
	largeSE := fill(grid, Position{X: 0, Y: 0}, largeSteps)
	largeNW := fill(grid, Position{X: size - 1, Y: 0}, largeSteps)
	largeSW := fill(grid, Position{X: size - 1, Y: size - 1}, largeSteps)

	sum := oddPoints*numOddGrids +
		evenPoints*numEvenGrids +
		cornerNPoints + cornerEPoints + cornerSPoints + cornerWPoints +
		(gridWidth+1)*(smallNE+smallSE+smallNW+smallSW) +
		(gridWidth)*(largeNE+largeSE+largeNW+largeSW)

	return strconv.Itoa(sum)
}

type QueueData struct {
	position       Position
	stepsRemaining int
}

func fill(grid Grid, start Position, steps int) int {
	ans := map[Position]bool{}
	seen := map[Position]bool{start: true}
	q := aq.New()
	q.Enqueue(QueueData{start, steps})
	for !q.Empty() {
		i, _ := q.Dequeue()
		qd := i.(QueueData)
		pos, s := qd.position, qd.stepsRemaining
		if s%2 == 0 {
			ans[pos] = true
		}
		if s == 0 {
			continue
		}
		possiblePositions := []Position{
			pos.Add(grids.North[int]()),
			pos.Add(grids.East[int]()),
			pos.Add(grids.South[int]()),
			pos.Add(grids.West[int]()),
		}
		for _, position := range possiblePositions {
			if _, ok := seen[position]; ok {
				continue
			}
			if !position.InBounds(0, 0, grid.Width()-1, grid.Height()-1) {
				continue
			}
			if grid.AtPos(position) == '#' {
				continue
			}
			// }
			seen[position] = true
			q.Enqueue(QueueData{position, s - 1})
		}
	}
	return len(ans)
}

func square(x int) int {
	return x * x
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
