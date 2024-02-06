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
	return strconv.Itoa(leastHeatLoss1(grid))
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := initializeGrid(lines)
	return strconv.Itoa(leastHeatLoss2(grid))
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

type State struct {
	pos      grids.Vector2D[int]
	dir      grids.Vector2D[int]
	steps    int
	heatLoss int
}

func ByHeatLoss(aI, bI interface{}) int {
	a := aI.(State)
	b := bI.(State)
	return a.heatLoss - b.heatLoss
}

func dirToMHLTIndex(direction grids.Vector2D[int]) int {
	switch direction {
	case grids.North[int]():
		return 0
	case grids.East[int]():
		return 1
	case grids.South[int]():
		return 2
	case grids.West[int]():
		return 3
	default:
		panic("Bad direction")
	}
}

func leastHeatLoss1(grid grids.Grid[int]) int {
	const MaxSteps = 3
	type Best = [4][MaxSteps + 1]int

	valid := func(pos grids.Vector2D[int]) bool {
		return pos.InBounds(0, 0, grid.Width()-1, grid.Height()-1)
	}
	lossAt := func(pos grids.Vector2D[int]) int {
		return grid.AtPos(pos)
	}

	mhlt := make([][]Best, grid.Height())
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

	q := priorityqueue.NewWith(ByHeatLoss)
	dest := grids.Vector2D[int]{X: grid.Width() - 1, Y: grid.Height() - 1}
	pushIfBetter := func(s State) {
		if mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] > s.heatLoss {
			mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] = s.heatLoss
			q.Enqueue(s)
		}
	}
	// known start states
	startPos := grids.Vector2D[int]{X: 0, Y: 0}
	pushIfBetter(State{pos: startPos, dir: grids.East[int](), steps: 0, heatLoss: 0})
	pushIfBetter(State{pos: startPos, dir: grids.South[int](), steps: 0, heatLoss: 0})

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

func leastHeatLoss2(grid grids.Grid[int]) int {
	const MaxSteps = 10
	type Best = [4][MaxSteps + 1]int

	valid := func(pos grids.Vector2D[int]) bool {
		return pos.InBounds(0, 0, grid.Width()-1, grid.Height()-1)
	}
	lossAt := func(pos grids.Vector2D[int]) int {
		return grid.AtPos(pos)
	}

	mhlt := make([][]Best, grid.Height())
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

	q := priorityqueue.NewWith(ByHeatLoss)
	dest := grids.Vector2D[int]{X: grid.Width() - 1, Y: grid.Height() - 1}
	pushIfBetter := func(s State) {
		if mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] > s.heatLoss {
			mhlt[s.pos.Y][s.pos.X][dirToMHLTIndex(s.dir)][s.steps] = s.heatLoss
			q.Enqueue(s)
		}
	}
	// known start states
	startPos := grids.Vector2D[int]{X: 0, Y: 0}
	pushIfBetter(State{pos: startPos, dir: grids.East[int](), steps: 0, heatLoss: 0})
	pushIfBetter(State{pos: startPos, dir: grids.South[int](), steps: 0, heatLoss: 0})

	for !q.Empty() {
		sI, _ := q.Peek()
		s := sI.(State)
		if s.pos == dest {
			return s.heatLoss
		}
		_, _ = q.Dequeue()

		forwardPos := s.pos.Add(s.dir)
		if s.steps < 10 && valid(forwardPos) {
			pushIfBetter(State{
				pos:      forwardPos,
				dir:      s.dir,
				steps:    s.steps + 1,
				heatLoss: s.heatLoss + lossAt(forwardPos),
			})
		}
		if s.steps >= 4 {
			leftPos := s.pos.Add(s.dir.RotateLeft())
			leftPos4 := s.pos.Add(s.dir.RotateLeft().ScalarProduct(4))
			if valid(leftPos4) {
				pushIfBetter(State{
					pos:      leftPos,
					dir:      s.dir.RotateLeft(),
					steps:    1,
					heatLoss: s.heatLoss + lossAt(leftPos),
				})
			}
			rightPos := s.pos.Add(s.dir.RotateRight())
			rightPos4 := s.pos.Add(s.dir.RotateRight().ScalarProduct(4))
			if valid(rightPos4) {
				pushIfBetter(State{
					pos:      rightPos,
					dir:      s.dir.RotateRight(),
					steps:    1,
					heatLoss: s.heatLoss + lossAt(rightPos),
				})
			}

		}
	}

	panic("Cannot get here")
}
