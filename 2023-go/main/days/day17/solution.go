package day17

import (
	"math"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/queues/priorityqueue"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := initializeGrid(lines)
	return strconv.Itoa(leastHeatLoss(grid))
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

func initializeGrid(lines []string) grids.Grid[int] {
	grid := grids.New[int](len(lines), len(lines[0]))
	for y, line := range lines {
		for x, i3 := range line {
			grid.SetAt(x, y, util.ConvertNumeric(string(i3)))
		}
	}
	return grid
}

const MaxSteps = 3

type State struct {
	pos      grids.Position
	dir      grids.Direction
	steps    int
	heatLoss int
}

func ByHeatLoss(aI, bI interface{}) int {
	a := aI.(State)
	b := bI.(State)
	return a.heatLoss - b.heatLoss
}

type Best = [4][MaxSteps + 1]int

func dirToMHLTIndex(direction grids.Direction) int {
	switch direction {
	case grids.North:
		return 0
	case grids.East:
		return 1
	case grids.South:
		return 2
	case grids.West:
		return 3
	default:
		panic("Bad direction")
	}
}

func leastHeatLoss(grid grids.Grid[int]) int {
	valid := func(pos grids.Position) bool {
		return pos.InBounds(0, 0, grid.Width()-1, grid.Height()-1)
	}
	lossAt := func(pos grids.Position) int {
		return grid.AtPos(pos)
	}

	mhlt := buildMinHeatLossTracker(grid)
	q := priorityqueue.NewWith(ByHeatLoss)
	dest := grids.Position{X: grid.Width() - 1, Y: grid.Height() - 1}
	pushIfBetter := func(s State) {
		if mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] > s.heatLoss {
			mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] = s.heatLoss
			q.Enqueue(s)
		}
	}
	// known start states
	startPos := grids.Position{X: 0, Y: 0}
	pushIfBetter(State{pos: startPos, dir: grids.East, steps: 0, heatLoss: 0})
	pushIfBetter(State{pos: startPos, dir: grids.South, steps: 0, heatLoss: 0})

	for !q.Empty() {
		sI, _ := q.Peek()
		s := sI.(State)
		if s.pos == dest {
			return s.heatLoss
		}
		_, _ = q.Dequeue()

		forwardPos := s.pos.Add(s.dir)
		if s.steps < 3 && valid(forwardPos) {
			pushIfBetter(State{
				pos:      forwardPos,
				dir:      s.dir,
				steps:    s.steps + 1,
				heatLoss: s.heatLoss + lossAt(forwardPos),
			})
		}
		leftPos := s.pos.Add(s.dir.RotateLeft())
		if valid(leftPos) {
			pushIfBetter(State{
				pos:      leftPos,
				dir:      s.dir.RotateLeft(),
				steps:    1,
				heatLoss: s.heatLoss + lossAt(leftPos),
			})
		}
		rightPos := s.pos.Add(s.dir.RotateRight())
		if valid(rightPos) {
			pushIfBetter(State{
				pos:      rightPos,
				dir:      s.dir.RotateRight(),
				steps:    1,
				heatLoss: s.heatLoss + lossAt(rightPos),
			})
		}
	}

	panic("Cannot get here")
}

func buildMinHeatLossTracker(grid grids.Grid[int]) (mhlt [][]Best) {
	mhlt = make([][]Best, grid.Height())
	for y := range mhlt {
		mhlt[y] = make([]Best, grid.Width())
		for x := range mhlt[y] {
			for b := range mhlt[y][x] {
				for a := range mhlt[y][x][b] {
					mhlt[y][x][b][a] = math.MaxInt64
				}
			}
		}
	}
	return
}
