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

	// Bridges of Königsberg puzzle.
	// Three "bridges" are enough to cut the graph into two.
	// Select a root node.
	// For every other node,
	// 		count the number of paths between that node and root without crossing
	// 		any edge more than once. If you have at least four paths, then that
	// 		other node *must* be on the same side of the three bridges (since
	// 		you can't cross a bridge twice). This cleanly divides the nodes
	// 		into two distinct groups.
	// We just need to count the paths once for the root node (any node will do --
	// there is nothing special about it) to every other node. Counting the paths
	// may be slow & expensive.

	return strconv.Itoa(0)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}
