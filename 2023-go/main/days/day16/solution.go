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
	beam := Beam{position: grids.Position{}, direction: grids.East}
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
		entryBeams = append(entryBeams, Beam{position: grids.Position{X: x, Y: 0}, direction: grids.South})
		entryBeams = append(entryBeams, Beam{position: grids.Position{X: x, Y: grid.Height() - 1}, direction: grids.North})
	}
	for y := 0; y < grid.Height(); y++ {
		entryBeams = append(entryBeams, Beam{position: grids.Position{X: 0, Y: y}, direction: grids.East})
		entryBeams = append(entryBeams, Beam{position: grids.Position{X: grid.Width() - 1, Y: y}, direction: grids.West})
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
	position  grids.Position
	direction grids.Direction
}

func initializeVisitedGrid(grid grids.Grid) grids.Grid {
	visited := grid.Clone()
	for y, row := range visited {
		for x := range row {
			visited.SetAt(x, y, '.')
		}
	}
	return visited
}

func countEnergizedTiles(grid grids.Grid) int {
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

func shootBeam(beam Beam, grid grids.Grid, visited grids.Grid, cache map[Beam]bool) {
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
		shootBeam(Beam{position: beam.position.Add(grids.Position(beam.direction)), direction: beam.direction}, grid, visited, cache)
	case '|':
		if beam.direction == grids.East || beam.direction == grids.West {
			// split N/S
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.North)), direction: grids.North}, grid, visited, cache)
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.South)), direction: grids.South}, grid, visited, cache)
		} else {
			// continue on
			shootBeam(Beam{position: beam.position.Add(grids.Position(beam.direction)), direction: beam.direction}, grid, visited, cache)
		}
	case '-':
		if beam.direction == grids.North || beam.direction == grids.South {
			// split E/W
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.East)), direction: grids.East}, grid, visited, cache)
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.West)), direction: grids.West}, grid, visited, cache)
		} else {
			// continue on
			shootBeam(Beam{position: beam.position.Add(grids.Position(beam.direction)), direction: beam.direction}, grid, visited, cache)
		}
	case '/':
		switch beam.direction {
		case grids.East:
			// turn north
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.North)), direction: grids.North}, grid, visited, cache)
		case grids.West:
			// turn south
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.South)), direction: grids.South}, grid, visited, cache)
		case grids.North:
			// turn east
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.East)), direction: grids.East}, grid, visited, cache)
		case grids.South:
			// turn west
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.West)), direction: grids.West}, grid, visited, cache)
		}
	case '\\':
		switch beam.direction {
		case grids.East:
			// turn south
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.South)), direction: grids.South}, grid, visited, cache)
		case grids.West:
			// turn north
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.North)), direction: grids.North}, grid, visited, cache)
		case grids.North:
			// turn west
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.West)), direction: grids.West}, grid, visited, cache)
		case grids.South:
			// turn east
			shootBeam(Beam{position: beam.position.Add(grids.Position(grids.East)), direction: grids.East}, grid, visited, cache)
		}
	default:
		panic("WTF")
	}
}
