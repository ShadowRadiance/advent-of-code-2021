package day15

import (
	"regexp"
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

	inputs := strings.Split(lines[0], ",")
	hashes := util.Transform(inputs, hasher)
	sum := util.Accumulate(hashes, func(total, next int) int { return total + next })

	return strconv.Itoa(sum)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	inputs := strings.Split(lines[0], ",")
	boxes := map[int][]Lens{}
	steps := util.Transform(inputs, func(s string) Step { return stepify(s) })
	for _, step := range steps {
		if step.operation == "-" {
			removeLens(step, boxes)
		} else if step.operation == "=" {
			addLens(step, boxes)
		}
	}

	sum := 0
	for boxNum, lenses := range boxes {
		sum += score(boxNum, lenses)
	}

	return strconv.Itoa(sum)
}

type Step struct {
	label       string
	boxNum      int
	operation   string
	focalLength int
}

type Lens struct {
	label       string
	focalLength int
}

func stepify(input string) Step {
	re := regexp.MustCompile(`(\w+)([-=])(\d*)`)
	matches := re.FindAllStringSubmatch(input, -1)

	step := Step{
		label:       matches[0][1],
		boxNum:      hasher(matches[0][1]),
		operation:   matches[0][2],
		focalLength: 0,
	}
	if step.operation == "=" {
		step.focalLength = util.ConvertNumeric(matches[0][3])
	}

	return step
}

func hasher(input string) int {
	currentValue := int32(0)
	for _, ascii := range input {
		currentValue += ascii
		currentValue *= 17
		currentValue %= 256
	}
	return int(currentValue)
}

func addLens(step Step, boxes map[int][]Lens) {
	newLens := Lens{label: step.label, focalLength: step.focalLength}

	existingIndex := slices.IndexFunc(boxes[step.boxNum], func(lens Lens) bool { return lens.label == newLens.label })
	if existingIndex == -1 {
		boxes[step.boxNum] = append(boxes[step.boxNum], newLens)
	} else {
		boxes[step.boxNum] = slices.Replace(boxes[step.boxNum], existingIndex, existingIndex+1, newLens)
	}
}

func removeLens(step Step, boxes map[int][]Lens) {
	boxes[step.boxNum] = slices.DeleteFunc(boxes[step.boxNum], func(lens Lens) bool { return lens.label == step.label })
}

func score(boxNum int, lenses []Lens) int {
	sum := 0
	for slot, lens := range lenses {
		sum += (boxNum + 1) * (slot + 1) * lens.focalLength
	}
	return sum
}
