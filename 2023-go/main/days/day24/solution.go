package day24

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type rect struct {
	x, y, h, w int
}

type Solution struct {
	testArea *rect
}

var part = 1

var live bool

func (s Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	if s.testArea == nil {
		live = true
		s.testArea = &rect{
			x: 200_000_000_000_000,
			y: 200_000_000_000_000,
			h: 200_000_000_000_000,
			w: 200_000_000_000_000,
		}
	}

	hailstones := parseHailstones(lines)

	intersecting := 0
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			if posf, ok := rayIntersect2D(hailstones[i], hailstones[j]); ok {
				if contains(s.testArea, posf) {
					if !live {
						fmt.Printf("Hailstones %d and %d intersect at (%f,%f)\n", i, j, posf.X, posf.Y)
					}
					intersecting += 1
				} else {
					if !live {
						fmt.Printf("Hailstones %d and %d intersect at (%f,%f) OUTSIDE TEST AREA\n", i, j, posf.X, posf.Y)
					}
				}
			} else {
				if !live {
					fmt.Printf("Hailstones %d and %d do not intersect\n", i, j)
				}
			}
		}
	}
	return strconv.Itoa(intersecting)
}

func (Solution) Part02(input string) string {
	part = 2

	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	live = len(lines) > 5

	hailstones := parseHailstones(lines)

	ourHailstone := solve(hailstones)
	// ourHailstone := solveUsingMatrix(hailstones)

	return strconv.FormatInt(
		ourHailstone.pos.X+
			ourHailstone.pos.Y+
			ourHailstone.pos.Z,
		10,
	)
}

// for any given hailstone,
// 		ABC = ABC.P + t(ABC.V),
// 	our stone,
// 		Q = Q.P + t(Q.V),
// 	should have a matching t/u/v value.
//
// With three stones (A,B,C), we have nine equations and nine unknowns (
// 		t, u, v, qpx, qpy, qpz, qvx, qvy, qvz
// ), assuming that any solution for three will work for all:
// 		APX + t(AVX) = qpx + t(qvx)
// 		APY + t(AVY) = qpy + t(qvy)
// 		APZ + t(AVZ) = qpz + t(qvz)
//
// 		BPX + u(BVX) = qpx + u(qvx)
// 		BPY + u(BVY) = qpy + u(qvy)
// 		BPZ + u(BVZ) = qpz + u(qvz)
//
// 		CPX + t(CVX) = qpx + v(qvx)
// 		CPY + t(CVY) = qpy + v(qvy)
// 		CPZ + t(CVZ) = qpz + v(qvz)
//
// We can eliminiate t, u, and v and end up with 6 equations with 6 unknowns (
// 		qpx, qpy, qpz, qvx, qvy, qvz
// ):
// 		(qpx - APX) / (AVX - qvx) = (qpy - APY) / (AVY - qvy) = (qpz - APZ) / (AVZ - qvz)
// 		(qpx - BPX) / (BVX - qvx) = (qpy - BPY) / (BVY - qvy) = (qpz - BPZ) / (BVZ - qvz)
// 		(qpx - CPX) / (CVX - qvx) = (qpy - CPY) / (CVY - qvy) = (qpz - CPZ) / (CVZ - qvz)
//
// Combining and simplifing the series of linear equations:
// 		[AVY - BVY]qpx - [AVX - BVX]qpy - [APY - BPY]qvx + [APX - BPX]qvy = (BPY * BVX - BPX * BVY) - (APY * AVX - APX * AVY)
// 		[AVY - CVY]qpx - [AVX - CVX]qpy - [APY - CPY]qvx + [APX - CPX]qvy = (CPY * CVX - CPX * CVY) - (APY * AVX - APX * AVY)
// 		[AVX - BVX]qpz - [AVZ - BVZ]qpx - [APX - BPX]qvz + [APZ - BPZ]qvx = (BPX * BVZ - BPZ * BVX) - (APX * AVZ - APZ * AVX)
// 		[AVX - CVX]qpz - [AVZ - CVZ]qpx - [APX - CPX]qvz + [APZ - CPZ]qvx = (CPX * CVZ - CPZ * CVX) - (APX * AVZ - APZ * AVX)
// 		[AVZ - BVZ]qpy - [AVY - BVY]qpz - [APZ - BPZ]qvy + [APY - BPY]qvz = (BPZ * BVY - BPY * BVZ) - (APZ * AVY - APY * AVZ)
// 		[AVZ - CVZ]qpy - [AVY - CVY]qpz - [APZ - CPZ]qvy + [APY - CPY]qvz = (CPZ * CVY - CPY * CVZ) - (APZ * AVY - APY * AVZ)
//
// Putting that in matrix notation (square_matrix_6x6 * column_unknowns_1x6 = column_constants_1x6)
//
//  | AVY - BVY,  BVX - AVX,          0,  BPY - APY,  APX - BPX,          0 |   | qpx |   | (BPY * BVX - BPX * BVY) - (APY * AVX - APX * AVY) |
//  | AVY - CVY,  CVX - AVX,          0,  CPY - APY,  APX - CPX,          0 |   | qpy |   | (CPY * CVX - CPX * CVY) - (APY * AVX - APX * AVY) |
//  | BVZ - AVZ,          0,  AVZ - BVX,  APZ - BPZ,          0,  BPX - APX | * | qpz | = | (BPX * BVZ - BPZ * BVX) - (APX * AVZ - APZ * AVX) |
//  | CVZ - AVZ,          0,  AVZ - CVX,  APZ - CPZ,          0,  CPX - APX |   | qvx |   | (CPX * CVZ - CPZ * CVX) - (APX * AVZ - APZ * AVX) |
//  |         0,  AVZ - BVZ,  BVY - AVY,          0,  BPZ - APZ,  APY - BPY |   | qvy |   | (BPZ * BVY - BPY * BVZ) - (APZ * AVY - APY * AVZ) |
//  |         0,  AVZ - CVZ,  CVY - AVY,          0,  CPZ - APZ,  APY - CPY |   | qvz |   | (CPZ * CVY - CPY * CVZ) - (APZ * AVY - APY * AVZ) |
//

func rat(a int64) *big.Rat {
	return big.NewRat(a, 1)
}

func solve(hailstones []Hailstone) (hs Hailstone) {
	a, b, c := chooseHailstones(hailstones)
	equations := setupEquations(a, b, c)
	equations = gaussianElimination(equations)

	// [
	// 	[1 0 0 0 0 0 | 24]
	// 	[0 1 0 0 0 0 | 13]
	// 	[0 0 1 0 0 0 | 10]
	// 	[0 0 0 1 0 0 | -3]
	// 	[0 0 0 0 1 0 | 1]
	// 	[0 0 0 0 0 1 | 2]
	// ]

	var f float64
	f, _ = equations[0][6].Float64()
	hs.pos.X = int64(f)
	f, _ = equations[1][6].Float64()
	hs.pos.Y = int64(f)
	f, _ = equations[2][6].Float64()
	hs.pos.Z = int64(f)
	f, _ = equations[3][6].Float64()
	hs.vel.X = int64(f)
	f, _ = equations[4][6].Float64()
	hs.vel.Y = int64(f)
	f, _ = equations[5][6].Float64()
	hs.vel.Z = int64(f)

	return
}

func chooseHailstones(hailstones []Hailstone) (a, b, c Hailstone) {
	a = hailstones[0]
	b = a
	c = a
	for i := 1; isParallel(a.vel, b.vel); i++ {
		b = hailstones[i]
	}
	for i := 2; isParallel(a.vel, c.vel) || isParallel(b.vel, c.vel); i++ {
		c = hailstones[i]
	}
	return
}

func setupEquations(a, b, c Hailstone) [][]*big.Rat {
	return [][]*big.Rat{
		{
			rat(0),
			rat(a.vel.Z - c.vel.Z),
			rat(c.vel.Y - a.vel.Y),
			rat(0),
			rat(c.pos.Z - a.pos.Z),
			rat(a.pos.Y - c.pos.Y),
			rat(a.pos.Y*a.vel.Z - a.pos.Z*a.vel.Y - c.pos.Y*c.vel.Z + c.pos.Z*c.vel.Y),
		},
		{
			rat(a.vel.Z - c.vel.Z),
			rat(0),
			rat(c.vel.X - a.vel.X),
			rat(c.pos.Z - a.pos.Z),
			rat(0),
			rat(a.pos.X - c.pos.X),
			rat(a.pos.X*a.vel.Z - a.pos.Z*a.vel.X - c.pos.X*c.vel.Z + c.pos.Z*c.vel.X),
		},
		{
			rat(c.vel.Y - a.vel.Y),
			rat(a.vel.X - c.vel.X),
			rat(0),
			rat(a.pos.Y - c.pos.Y),
			rat(c.pos.X - a.pos.X),
			rat(0),
			rat(a.pos.Y*a.vel.X - a.pos.X*a.vel.Y - c.pos.Y*c.vel.X + c.pos.X*c.vel.Y),
		},
		{
			rat(0),
			rat(b.vel.Z - c.vel.Z),
			rat(c.vel.Y - b.vel.Y),
			rat(0),
			rat(c.pos.Z - b.pos.Z),
			rat(b.pos.Y - c.pos.Y),
			rat(b.pos.Y*b.vel.Z - b.pos.Z*b.vel.Y - c.pos.Y*c.vel.Z + c.pos.Z*c.vel.Y),
		},
		{
			rat(b.vel.Z - c.vel.Z),
			rat(0),
			rat(c.vel.X - b.vel.X),
			rat(c.pos.Z - b.pos.Z),
			rat(0),
			rat(b.pos.X - c.pos.X),
			rat(b.pos.X*b.vel.Z - b.pos.Z*b.vel.X - c.pos.X*c.vel.Z + c.pos.Z*c.vel.X),
		},
		{
			rat(c.vel.Y - b.vel.Y),
			rat(b.vel.X - c.vel.X),
			rat(0),
			rat(b.pos.Y - c.pos.Y),
			rat(c.pos.X - b.pos.X),
			rat(0),
			rat(b.pos.Y*b.vel.X - b.pos.X*b.vel.Y - c.pos.Y*c.vel.X + c.pos.X*c.vel.Y),
		},
	}
}

func gaussianElimination(equations [][]*big.Rat) [][]*big.Rat {
	dump(equations)

	// Iterate diagonally from top left, to turn matrix into reduced row echelon form
	for i := 0; i < len(equations); i++ {
		// Find non-zero item in current column, from current row or after
		var nonZeroRowIndex = -1
		for j := i; j < len(equations); j++ {
			if equations[j][i].Cmp(rat(0)) != 0 {
				nonZeroRowIndex = j
			}
		}

		// Swap current row with first non-zero row
		if nonZeroRowIndex != i {
			equations[i], equations[nonZeroRowIndex] = equations[nonZeroRowIndex], equations[i]
		}

		// Divide row by value at current pos, to turn value into 1
		currVal := equations[i][i]
		equations[i][i] = rat(1)
		for c := i + 1; c < len(equations[i]); c++ {
			equations[i][c].Quo(equations[i][c], currVal)
		}

		// Subtract multiple of current row from lower rows, to turn column below current item to 0
		for r := i + 1; r < len(equations); r++ {
			multiple := equations[r][i]
			equations[r][i] = rat(0)
			if multiple.Cmp(rat(0)) != 0 {
				for c := i + 1; c < len(equations[r]); c++ {
					res := rat(0)
					res.Mul(equations[i][c], multiple)
					equations[r][c].Sub(equations[r][c], res)
				}
			}
		}
		dump(equations)
	}

	// Iterate diagonally from bottom right, to turn matrix (except last column) into unit matrix.
	lastColumn := len(equations[0]) - 1
	for rc := len(equations) - 1; rc >= 0; rc-- {
		for r := 0; r < rc; r++ {
			res := rat(0)
			res.Mul(equations[rc][lastColumn], equations[r][rc])
			equations[r][lastColumn].Sub(equations[r][lastColumn], res)
			equations[r][rc] = rat(0)
		}
		dump(equations)
	}

	return equations
}

func dump(equations [][]*big.Rat) {
	if !live {
		fmt.Println(equations)
	}
}

func solveUsingMatrix(hailstones []Hailstone) Hailstone {
	a, b, c := chooseHailstones(hailstones)
	fmt.Println("Input Hailstones: ", a, b, c)

	squareMatrix := [][]int64{
		// 6 rows, 6 columns
		{a.vel.Y - b.vel.Y, b.vel.X - a.vel.X, 0, b.pos.Y - a.pos.Y, a.pos.X - b.pos.X, 0},
		{a.vel.Y - c.vel.Y, c.vel.X - a.vel.X, 0, c.pos.Y - a.pos.Y, a.pos.X - c.pos.X, 0},
		{b.vel.Z - a.vel.Z, 0, a.vel.Z - b.vel.X, a.pos.Z - b.pos.Z, 0, b.pos.X - a.pos.X},
		{c.vel.Z - a.vel.Z, 0, a.vel.Z - c.vel.X, a.pos.Z - c.pos.Z, 0, c.pos.X - a.pos.X},
		{0, a.vel.Z - b.vel.Z, b.vel.Y - a.vel.Y, 0, b.pos.Z - a.pos.Z, a.pos.Y - b.pos.Y},
		{0, a.vel.Z - c.vel.Z, c.vel.Y - a.vel.Y, 0, c.pos.Z - a.pos.Z, a.pos.Y - c.pos.Y},
	}

	fmt.Println("Initial Matrix: ", squareMatrix)

	columnConstants := [][]int64{
		// 6 rows, 1 column
		{(b.pos.Y*b.vel.X - b.pos.X*b.vel.Y) - (a.pos.Y*a.vel.X - a.pos.X*a.vel.Y)},
		{(c.pos.Y*c.vel.X - c.pos.X*c.vel.Y) - (a.pos.Y*a.vel.X - a.pos.X*a.vel.Y)},
		{(b.pos.X*b.vel.Z - b.pos.Z*b.vel.X) - (a.pos.X*a.vel.Z - a.pos.Z*a.vel.X)},
		{(c.pos.X*c.vel.Z - c.pos.Z*c.vel.X) - (a.pos.X*a.vel.Z - a.pos.Z*a.vel.X)},
		{(b.pos.Z*b.vel.Y - b.pos.Y*b.vel.Z) - (a.pos.Z*a.vel.Y - a.pos.Y*a.vel.Z)},
		{(c.pos.Z*c.vel.Y - c.pos.Y*c.vel.Z) - (a.pos.Z*a.vel.Y - a.pos.Y*a.vel.Z)},
	}

	fmt.Println("Initial Constants: ", columnConstants)

	inverseMatrix := invert(squareMatrix) // 6 rows, 6 columns

	fmt.Println("Inverted Matrix: ", inverseMatrix)

	columnUnknowns := multiply(inverseMatrix, columnConstants) // 6 rows, 1 column

	fmt.Println("Unknowns: ", columnUnknowns)

	hs := Hailstone{
		pos: Pos{
			X: columnUnknowns[0][0],
			Y: columnUnknowns[1][0],
			Z: columnUnknowns[2][0],
		},
		vel: Pos{
			X: columnUnknowns[3][0],
			Y: columnUnknowns[4][0],
			Z: columnUnknowns[5][0],
		},
	}

	fmt.Println("Hailstone: ", hs)
	return hs
}

type Pos = grids.Vector3D[int64]
type PosF = grids.Vector3D[float64]
type Vel = grids.Vector3D[int64]
type VelF = grids.Vector3D[float64]

// Hailstone is a ray starting at position pos and moving in direction vel
type Hailstone struct {
	pos Pos
	vel Vel
}

func parseHailstones(lines []string) (hss []Hailstone) {
	for _, line := range lines {
		hss = append(hss, parseHailstone(line))
	}
	return
}

func parseHailstone(line string) (hs Hailstone) {
	// 240883930774627, 293767063383987, 385416738115181 @ 85, 41, -232
	posVel := strings.Split(line, " @ ")
	posParts := strings.Split(strings.TrimSpace(posVel[0]), ", ")
	velParts := strings.Split(strings.TrimSpace(posVel[1]), ", ")
	xPos, _ := strconv.ParseInt(strings.TrimSpace(posParts[0]), 10, 64)
	yPos, _ := strconv.ParseInt(strings.TrimSpace(posParts[1]), 10, 64)
	var zPos int64 = 0
	if part == 2 {
		zPos, _ = strconv.ParseInt(strings.TrimSpace(posParts[2]), 10, 64)
	}
	hs.pos = Pos{X: xPos, Y: yPos, Z: zPos}
	xVel, _ := strconv.ParseInt(strings.TrimSpace(velParts[0]), 10, 64)
	yVel, _ := strconv.ParseInt(strings.TrimSpace(velParts[1]), 10, 64)
	var zVel int64 = 0
	if part == 2 {
		zVel, _ = strconv.ParseInt(strings.TrimSpace(velParts[2]), 10, 64)
	}
	hs.vel = Vel{X: xVel, Y: yVel, Z: zVel}

	return
}

func rayIntersect2D(a, b Hailstone) (PosF, bool) {
	aPosF := PosF{X: float64(a.pos.X), Y: float64(a.pos.Y)}
	bPosF := PosF{X: float64(b.pos.X), Y: float64(b.pos.Y)}
	aVelF := VelF{X: float64(a.vel.X), Y: float64(a.vel.Y)}
	bVelF := VelF{X: float64(b.vel.X), Y: float64(b.vel.Y)}

	if aPosF == bPosF {
		return aPosF, true
	}

	dx := bPosF.X - aPosF.X
	dy := bPosF.Y - aPosF.Y
	det := (bVelF.X * aVelF.Y) - (bVelF.Y * aVelF.X)

	if det != 0 {
		u := ((dy * bVelF.X) - (dx * bVelF.Y)) / det
		v := ((dy * aVelF.X) - (dx * aVelF.Y)) / det
		if u >= 0.0 && v >= 0.0 {
			return aPosF.Add(aVelF.ScalarProduct(u)), true
		}
	}

	return PosF{}, false
}

func contains(r *rect, p PosF) bool {
	return p.X >= float64(r.x) && p.X <= float64(r.x+r.w) &&
		p.Y >= float64(r.y) && p.Y <= float64(r.y+r.h)
}

func isParallel(vel1, vel2 Vel) bool {
	ratioX := float64(vel1.X) / float64(vel2.X)
	ratioY := float64(vel1.Y) / float64(vel2.Y)
	ratioZ := float64(vel1.Z) / float64(vel2.Z)
	return ratioX == ratioY && ratioX == ratioZ
}

func invert(matrix [][]int64) [][]int64 {
	rows := len(matrix)
	if rows == 0 {
		panic("multiplicand has no rows")
	}
	cols := len(matrix[0])
	if cols == 0 {
		panic("multiplicand has no columns")
	}

	if rows != cols {
		panic("Cannot invert a non-square matrix")
	}

	result := make([][]int64, rows)
	for i := range result {
		result[i] = make([]int64, cols)
	}

	// Actually invert the 6x6 matrix... okay... Wolfram?
	// matrix {{-408,-150,0,173431832680772,92119073828900,0},{180,-29,0,224319115594086,51058981961611,0},{-54,0,122,104016724747858,0,-92119073828900},{167,0,1,149438480487458,0,-51058981961611},{0,54,408,0,-104016724747858,-173431832680772},{0,-167,-180,0,-149438480487458,-224319115594086}}
	// consts {{-156443739915072178},{42613167847055594},{25512231958296218},{57193229613436778},{123047441338600000},{-97660155204604219}}
	// inverse {{0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0.43477, 0.675001}, {0, 0, 0, 0, 0, 0}, {0, 0, 0, 0, -0.43477, 0.324999}}
	// too many empty rows!

	return result
}

func multiply(matrix1 [][]int64, matrix2 [][]int64) [][]int64 {
	rows1 := len(matrix1)
	if rows1 == 0 {
		panic("multiplicand has no rows")
	}
	cols1 := len(matrix1[0])
	if cols1 == 0 {
		panic("multiplicand has no columns")
	}
	rows2 := len(matrix2)
	if rows2 == 0 {
		panic("multiplicand has no rows")
	}
	cols2 := len(matrix2[0])
	if cols2 == 0 {
		panic("multiplicand has no columns")
	}

	if cols1 != rows2 {
		panic("matrices must be in the form NxM * MxL")
	}

	// resultant matrix will be NxL
	result := make([][]int64, rows1)
	for r, _ := range result {
		result[r] = make([]int64, cols2)
		for c, _ := range result[r] {
			for index := 0; index < cols1; index++ {
				result[r][c] += matrix1[r][index] * matrix2[index][c]
			}
		}
	}
	return result
}

// 505689961281052 too low
// 936849059171706 wrong
// 957615498525366 wrong
