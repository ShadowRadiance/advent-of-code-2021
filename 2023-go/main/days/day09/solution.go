package day09

import (
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	histories := parseHistories(lines)

	for i := range histories {
		history := &histories[i]
		history.addRightValue()
	}

	values := make([]int, 0, len(histories))
	for i := range histories {
		history := &histories[i]
		last := len(history.steps[0]) - 1
		values = append(values, history.steps[0][last])
	}

	return strconv.Itoa(
		util.Accumulate(values, func(a, b int) int { return a + b }))
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.Trim(input, "\n"), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	histories := parseHistories(lines)

	for i := range histories {
		history := &histories[i]
		history.addLeftValue()
	}

	leftValues := make([]int, 0, len(histories))
	for i := range histories {
		history := &histories[i]
		leftValues = append(leftValues, history.steps[0][0])
	}

	return strconv.Itoa(
		util.Accumulate(leftValues, func(a, b int) int { return a + b }))
}

type History struct {
	initial []int
	steps   [][]int
}

func isZero(value int) bool { return value == 0 }

func (h *History) addRightValue() {
	h.buildSteps()

	// build back the finals
	last := len(h.steps) - 1
	for i := last; i >= 0; i-- {
		if util.All(h.steps[i], isZero) {
			h.steps[i] = append(h.steps[i], 0)
		} else {
			lastInStep := len(h.steps[i]) - 1
			h.steps[i] = append(h.steps[i],
				h.steps[i][lastInStep]+h.steps[i+1][lastInStep])
		}
	}

	// fmt.Printf("%+v\n", h)
	// fmt.Printf("---\n\n")

	return
}

func (h *History) buildSteps() {
	h.steps = [][]int{h.initial}
	// while the last step is not all zeroes
	last := len(h.steps) - 1
	for !util.All(h.steps[last], isZero) {
		// add a row with the differences
		h.steps = append(h.steps, buildDifferenceRow(h.steps[last]))
		last++
		if len(h.steps[last]) == 1 && h.steps[last][0] != 0 {
			panic("WAT!")
		}
	}
}

func (h *History) addLeftValue() {
	if len(h.steps) == 0 {
		h.buildSteps()
	}

	// fmt.Printf("%+v\n", h)

	for i, step := range h.steps {
		h.steps[i] = prepend(step, 0)
	}

	// fmt.Printf("%+v\n", h)

	// build back the prefixes
	last := len(h.steps) - 1
	for i := last - 1; i >= 0; i-- {
		h.steps[i][0] = h.steps[i][1] - h.steps[i+1][0]
	}

	// fmt.Printf("%+v\n", h)
	// fmt.Printf("---\n\n")

	return
}

func parseHistories(lines []string) (histories []History) {
	for _, line := range lines {
		histories = append(histories, parseHistory(line))
	}
	return
}

func parseHistory(line string) (history History) {
	history.initial = util.MapStringsToIntegers(strings.Split(line, " "))
	return
}

func buildDifferenceRow(ints []int) (result []int) {
	for i := 1; i < len(ints); i++ {
		result = append(result, ints[i]-ints[i-1])
	}
	return
}

func prepend(ints []int, value int) []int {
	ints = append(ints, value)
	for i := len(ints) - 2; i >= 0; i-- {
		ints[i+1] = ints[i]
	}
	ints[0] = value
	return ints
}
