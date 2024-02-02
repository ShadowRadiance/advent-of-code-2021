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
	slices.SortFunc(bricks, func(a, b *Brick) int {
		if a.loZ() == b.loZ() {
			return cmp.Compare(a.hiZ(), b.hiZ())
		}
		return cmp.Compare(a.loZ(), b.loZ())
	})
	bricks = settle(bricks)

	sum := 0
	for _, brick := range bricks {
		if canDisintegrate(brick) {
			sum++
		}
	}
	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	bricks := parseBricks(lines)
	slices.SortFunc(bricks, func(a, b *Brick) int {
		if a.loZ() == b.loZ() {
			return cmp.Compare(a.hiZ(), b.hiZ())
		}
		return cmp.Compare(a.loZ(), b.loZ())
	})
	bricks = settle(bricks)

	sum := 0
	for _, brick := range bricks {
		ccr := countChainReaction(brick)
		sum += ccr
	}
	return strconv.Itoa(sum)
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
func (b *Brick) fall() {
	// fmt.Printf("Dropping brick %s\n", b.name)
	b.z1 -= 1
	b.z2 -= 1
}

func (b *Brick) String() string {
	return fmt.Sprintf(
		"'%s' (%d,%d,%d-%d,%d,%d) supporting: [%s] supportedBy: [%s]",
		b.name, b.x1, b.y1, b.z1, b.x2, b.y2, b.z2,
		strings.Join(util.Transform(b.supporting, func(sB *Brick) string { return sB.name }), " "),
		strings.Join(util.Transform(b.supportedBy, func(sB *Brick) string { return sB.name }), " "))
}

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

func settle(bricks []*Brick) (settledBricks []*Brick) {
	settledBricks = make([]*Brick, 0)

	for _, brick := range bricks {
		testBrick := *brick
		testBrick.name = "TEST-" + brick.name
		testBrick.fall()
		overlapping := testBrick.overlapping(settledBricks)
		for len(overlapping) == 0 {
			if testBrick.loZ() < 1 {
				break
			}

			// copy fall onto original brick
			brick.fall()

			// set up next loop
			testBrick.fall()
			overlapping = testBrick.overlapping(settledBricks)
		}
		brick.supportedBy = overlapping

		for _, supporter := range brick.supportedBy {
			supporter.supporting = append(supporter.supporting, brick)
		}
		settledBricks = append(settledBricks, brick)
	}

	return
}

func canDisintegrate(brick *Brick) bool {
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

func countChainReaction(brick *Brick) int {
	// for a given brick (brick),
	// we want to find the bricks that *would fall* if it stopped supporting them
	// - ie. the ones it supports that are NOT supported by another brick
	// along with the chain reaction for each of those bricks stopping supporting their dependents, etc

	fallingDependents := make([]string, 0)

	maybeDependents := make([]*Brick, 0)
	maybeDependents = slices.Clone(brick.supporting)

	for len(maybeDependents) > 0 {
		checkers := slices.Clone(maybeDependents)
		maybeDependents = make([]*Brick, 0)

		for _, checkDependent := range checkers {
			if isDependent(checkDependent, brick, fallingDependents) {
				if !slices.Contains(fallingDependents, checkDependent.name) {
					fallingDependents = append(fallingDependents, checkDependent.name)
				}
				for _, supported := range checkDependent.supporting {
					if !slices.Contains(maybeDependents, supported) {
						maybeDependents = append(maybeDependents, supported)
					}
				}
			}
		}
	}

	return len(fallingDependents)
}

func isDependent(check *Brick, fallenParent *Brick, fallenDependencies []string) bool {
	// if the check brick is supported by anything *except* the fallen items, return false
	for _, supporter := range check.supportedBy {
		if fallenParent.name == supporter.name || slices.Contains(fallenDependencies, supporter.name) {
			continue
		}
		return false
	}
	// otherwise nothing is supporting this brick - it is dependent
	return true
}
