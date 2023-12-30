package day21

import (
	"strings"
)

type Solution struct {
	steps int
}

func (s Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	if s.steps == 0 {
		s.steps = 64
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
