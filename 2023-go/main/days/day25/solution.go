package day25

import (
	"fmt"
	"strconv"
	"strings"
)

type Graph = map[string][]string

func parseLines(lines []string) (lookup Graph) {
	lookup = Graph{}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rootName := parts[0]
		connectedNames := strings.Split(strings.TrimSpace(parts[1]), " ")

		if _, ok := lookup[rootName]; !ok {
			lookup[rootName] = connectedNames
		} else {
			lookup[rootName] = append(lookup[rootName], connectedNames...)
		}

		for _, connectedName := range connectedNames {
			if _, ok := lookup[connectedName]; !ok {
				lookup[connectedName] = []string{rootName}
			} else {
				lookup[connectedName] = append(lookup[connectedName], rootName)
			}
		}
	}
	return
}

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	lookup := parseLines(lines)
	fmt.Println(lookup)
	return strconv.Itoa(0)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}
