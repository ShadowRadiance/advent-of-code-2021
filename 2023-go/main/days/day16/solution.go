package day16

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := grids.NewGrid(lines)
	visited := initializeVisitedGrid(grid)
	cache := map[Beam]bool{}
	beam := Beam{position: grids.Position[int]{}, direction: grids.East[int]()}
	shootBeam(beam, grid, visited, cache)

	numberOfEnergizedTiles := countEnergizedTiles(visited)
	return strconv.Itoa(numberOfEnergizedTiles)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := grids.NewGrid(lines)

	highestNumberOfEnergizedTiles := 0

	entryBeams := make([]Beam, 0)
	for x := 0; x < grid.Width(); x++ {
		entryBeams = append(entryBeams, Beam{position: grids.Position[int]{X: x, Y: 0}, direction: grids.South[int]()})
		entryBeams = append(entryBeams, Beam{position: grids.Position[int]{X: x, Y: grid.Height() - 1}, direction: grids.North[int]()})
	}
	for y := 0; y < grid.Height(); y++ {
		entryBeams = append(entryBeams, Beam{position: grids.Position[int]{X: 0, Y: y}, direction: grids.East[int]()})
		entryBeams = append(entryBeams, Beam{position: grids.Position[int]{X: grid.Width() - 1, Y: y}, direction: grids.West[int]()})
	}

	for _, beam := range entryBeams {
		visited := initializeVisitedGrid(grid)
		cache := map[Beam]bool{}
		shootBeam(beam, grid, visited, cache)
		numberOfEnergizedTiles := countEnergizedTiles(visited)
		if numberOfEnergizedTiles > highestNumberOfEnergizedTiles {
			highestNumberOfEnergizedTiles = numberOfEnergizedTiles
		}
	}

	return strconv.Itoa(highestNumberOfEnergizedTiles)
}

type Beam struct {
	position  grids.Position[int]
	direction grids.Direction[int]
}

func initializeVisitedGrid(grid grids.Grid[rune]) grids.Grid[rune] {
	visited := grid.Clone()
	for y, row := range visited {
		for x := range row {
			visited.SetAt(x, y, '.')
		}
	}
	return visited
}

func countEnergizedTiles(grid grids.Grid[rune]) int {
	numberOfEnergizedTiles := 0
	for _, row := range grid {
		for _, char := range row {
			if char == '#' {
				numberOfEnergizedTiles++
			}
		}
	}
	return numberOfEnergizedTiles
}

func shootBeam(beam Beam, grid grids.Grid[rune], visited grids.Grid[rune], cache map[Beam]bool) {
	if remembered, ok := cache[beam]; ok && remembered {
		// we hit a cycle and we're done with this beam
		return
	}
	cache[beam] = true
	if !beam.position.InBounds(0, 0, grid.Width()-1, grid.Height()-1) {
		// we exited the map, we're done with this beam
		return
	}

	visited.SetAtPos(beam.position, '#')
	switch grid.AtPos(beam.position) {
	case '.':
		// continue on
		shootBeam(Beam{position: beam.position.Add(beam.direction), direction: beam.direction}, grid, visited, cache)
	case '|':
		if beam.direction == grids.East[int]() || beam.direction == grids.West[int]() {
			// split N/S
			shootBeam(Beam{position: beam.position.Add(grids.North[int]()), direction: grids.North[int]()}, grid, visited, cache)
			shootBeam(Beam{position: beam.position.Add(grids.South[int]()), direction: grids.South[int]()}, grid, visited, cache)
		} else {
			// continue on
			shootBeam(Beam{position: beam.position.Add(beam.direction), direction: beam.direction}, grid, visited, cache)
		}
	case '-':
		if beam.direction == grids.North[int]() || beam.direction == grids.South[int]() {
			// split E/W
			shootBeam(Beam{position: beam.position.Add(grids.East[int]()), direction: grids.East[int]()}, grid, visited, cache)
			shootBeam(Beam{position: beam.position.Add(grids.West[int]()), direction: grids.West[int]()}, grid, visited, cache)
		} else {
			// continue on
			shootBeam(Beam{position: beam.position.Add(beam.direction), direction: beam.direction}, grid, visited, cache)
		}
	case '/':
		switch beam.direction {
		case grids.East[int]():
			// turn north
			shootBeam(Beam{position: beam.position.Add(grids.North[int]()), direction: grids.North[int]()}, grid, visited, cache)
		case grids.West[int]():
			// turn south
			shootBeam(Beam{position: beam.position.Add(grids.South[int]()), direction: grids.South[int]()}, grid, visited, cache)
		case grids.North[int]():
			// turn east
			shootBeam(Beam{position: beam.position.Add(grids.East[int]()), direction: grids.East[int]()}, grid, visited, cache)
		case grids.South[int]():
			// turn west
			shootBeam(Beam{position: beam.position.Add(grids.West[int]()), direction: grids.West[int]()}, grid, visited, cache)
		}
	case '\\':
		switch beam.direction {
		case grids.East[int]():
			// turn south
			shootBeam(Beam{position: beam.position.Add(grids.South[int]()), direction: grids.South[int]()}, grid, visited, cache)
		case grids.West[int]():
			// turn north
			shootBeam(Beam{position: beam.position.Add(grids.North[int]()), direction: grids.North[int]()}, grid, visited, cache)
		case grids.North[int]():
			// turn west
			shootBeam(Beam{position: beam.position.Add(grids.West[int]()), direction: grids.West[int]()}, grid, visited, cache)
		case grids.South[int]():
			// turn east
			shootBeam(Beam{position: beam.position.Add(grids.East[int]()), direction: grids.East[int]()}, grid, visited, cache)
		}
	default:
		panic("WTF")
	}
}
