package day15

import (
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

	return "PENDING"
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
