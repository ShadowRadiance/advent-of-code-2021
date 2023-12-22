package day18

import (
	"math"
	"regexp"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return solution(parseDigPlan(lines))
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return solution(parseDigPlanV2(lines))
}

func solution(digPlan DigPlan) string {
	trench := digPlan.digTrenchV2(Pos{})

	area := shoelace(trench)
	boundary := trenchLength(trench)
	interior := interiorPoints(area, boundary)
	return strconv.Itoa(interior + boundary)

}

type Pos = grids.Position[float64]
type Dir = grids.Direction[float64]

type DigInstruction struct {
	dir      Dir
	distance float64
}
type DigPlan []DigInstruction

func (d DigPlan) digTrench(start Pos) (list []Pos) {
	list = append(list, start)
	for _, inst := range d {
		for i := 0; i < int(inst.distance); i++ {
			start = start.Add(inst.dir)
			list = append(list, start)
		}
	}
	return
}

func (d DigPlan) digTrenchV2(start Pos) (list []Pos) {
	vertex := start
	list = append(list, vertex)
	for _, inst := range d {
		vertex = vertex.Add(inst.dir.ScalarProduct(inst.distance))
		list = append(list, vertex)
	}
	return
}

func parseDigPlan(lines []string) (plan DigPlan) {
	for _, line := range lines {
		plan = append(plan, parseDigInstruction(line))
	}
	return
}

var (
	North = grids.North[float64]()
	East  = grids.East[float64]()
	West  = grids.West[float64]()
	South = grids.South[float64]()
)

func parseDigInstruction(line string) (di DigInstruction) {
	re := regexp.MustCompile(`([A-Z]) (\d+) \(#[0-9A-Fa-f]{6}\)`)

	matches := re.FindAllStringSubmatch(line, -1)

	switch matches[0][1] {
	case "U":
		di.dir = North
	case "R":
		di.dir = East
	case "D":
		di.dir = South
	case "L":
		di.dir = West
	}

	di.distance = float64(util.ConvertNumeric(matches[0][2]))

	return
}

func parseDigPlanV2(lines []string) (plan DigPlan) {
	for _, line := range lines {
		plan = append(plan, parseDigInstructionV2(line))
	}
	return
}

func parseDigInstructionV2(line string) (di DigInstruction) {
	re := regexp.MustCompile(`[A-Z] \d+ \(#([0-9A-Fa-f]{5})([0-9A-Fa-f])\)`)

	matches := re.FindAllStringSubmatch(line, -1)

	switch matches[0][2] {
	case "0":
		di.dir = East
	case "1":
		di.dir = South
	case "2":
		di.dir = West
	case "3":
		di.dir = North
	}

	i64, _ := strconv.ParseInt(matches[0][1], 16, 64)
	di.distance = float64(i64)

	return
}

func shoelace(points []Pos) float64 {
	area := 0.0

	points = append(points, points[0])
	for i := 1; i < len(points); i++ {
		area += (points[i-1].X * points[i].Y) - (points[i-1].Y * points[i].X)
	}

	return math.Abs(area) / 2
}

func interiorPoints(area float64, numBoundaryPoints int) int {
	return int(area - float64(numBoundaryPoints)/2.0 + 1.0)
}

func trenchLength(trench []Pos) int {
	length := 0.0
	for i := 1; i < len(trench); i++ {
		if trench[i].X == trench[i-1].X {
			length += math.Abs(trench[i].Y - trench[i-1].Y)
		} else {
			length += math.Abs(trench[i].X - trench[i-1].X)
		}
	}
	return int(length)
}
