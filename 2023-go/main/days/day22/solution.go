package day22

import (
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	bricks := parseBricks(lines)
	slices.SortFunc(bricks, func(a, b *Brick) int { return cmp.Compare(a.z1, b.z1) })
	settle(bricks)
	sum := 0
	for _, brick := range bricks {
		if canDisintegrate(brick, bricks) {
			sum++
		}
	}

	for _, brick := range bricks {
		fmt.Printf("%+v\n", brick)
	}
	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

type Brick struct {
	name        string
	x1, y1, z1  int
	x2, y2, z2  int
	supporting  []*Brick
	supportedBy []*Brick
}

func (b *Brick) loX() int { return b.x1 }
func (b *Brick) hiX() int { return b.x2 }
func (b *Brick) loY() int { return b.y1 }
func (b *Brick) hiY() int { return b.y2 }
func (b *Brick) loZ() int { return b.z1 }
func (b *Brick) hiZ() int { return b.z2 }

func (b *Brick) overlapping(bricks []*Brick) (supporters []*Brick) {
	// reduce the list of what we have to check for overlapping!
	for _, other := range bricks {
		if b.name == other.name {
			continue
		}
		if b.overlaps(other) {
			supporters = append(supporters, other)
		}
	}
	return
}

func (b *Brick) overlaps(other *Brick) bool {
	return b.loX() <= other.hiX() && b.hiX() >= other.loX() &&
		b.loY() <= other.hiY() && b.hiY() >= other.loY() &&
		b.loZ() <= other.hiZ() && b.hiZ() >= other.loZ()
}

func parseBricks(lines []string) (bricks []*Brick) {
	for i, line := range lines {
		name := fmt.Sprintf("A%.4x", i)
		bricks = append(bricks, parseBrick(line, name))
	}
	return
}

func parseBrick(line string, name string) (brick *Brick) {
	brick = &Brick{name: name, supporting: make([]*Brick, 0), supportedBy: make([]*Brick, 0)}
	for i, s := range strings.Split(line, "~") {
		coords := util.Transform(strings.Split(s, ","), func(item string) int { return util.ConvertNumeric(item) })
		if i == 0 {
			brick.x1, brick.y1, brick.z1 = coords[0], coords[1], coords[2]
		} else {
			brick.x2, brick.y2, brick.z2 = coords[0], coords[1], coords[2]
		}
	}
	return
}

func settle(bricks []*Brick) {
	for _, brick := range bricks {
		fmt.Printf("Dropping brick %s\n", brick.name)

		// for each brick, move it down (z=z-1) if it won't overlap another brick
		if brick.loZ() == 1 {
			continue
		}

		testBrick := *brick
		testBrick.z1, testBrick.z2 = testBrick.z1-1, testBrick.z2-1
		// reduce the list of what we have to check for overlapping!
		overlapping := testBrick.overlapping(bricks)
		for len(overlapping) == 0 {
			brick.z1, brick.z2 = testBrick.z1, testBrick.z2
			testBrick.z1, testBrick.z2 = testBrick.z1-1, testBrick.z2-1
			// reduce the list of what we have to check for overlapping!
			overlapping = testBrick.overlapping(bricks)
		}
		brick.supportedBy = overlapping

		for _, supporter := range brick.supportedBy {
			supporter.supporting = append(supporter.supporting, brick)
		}
	}

	return
}

func canDisintegrate(brick *Brick, bricks []*Brick) bool {
	if len(brick.supporting) == 0 {
		return true
	}

	for _, supported := range brick.supporting {
		if len(supported.supportedBy) == 1 {
			return false
		}
	}

	return true
}
