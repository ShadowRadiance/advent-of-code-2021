package day14

import (
	"strings"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}
