package day08

import (
	"cmp"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	instructions, nodes := parseInput(lines)

	moves := 0
	index, _ := slices.BinarySearchFunc(
		nodes,
		&Node{"AAA", "", "", [2]*Node{}},
		func(a *Node, b *Node) int {
			return cmp.Compare(a.Name, b.Name)
		})
	currentNode := nodes[index]
	for currentNode.Name != "ZZZ" {
		currentNode = processInstruction(currentNode, instructions, moves)
		moves++
	}

	return strconv.Itoa(moves)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(input, "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	instructions, nodes := parseInput(lines)

	currentNodes := make([]*Node, 0)
	for _, node := range nodes {
		if strings.HasSuffix(node.Name, "A") {
			currentNodes = append(currentNodes, node)
		}
	}

	// // brute force
	// moves := 0
	// for !allEndInZ(currentNodes) {
	// 	for i, node := range currentNodes {
	// 		currentNodes[i] = processInstruction(node, instructions, moves)
	// 	}
	// 	moves++
	// }
	// return strconv.Itoa(moves)

	// LCM happens to work for our datasets because of cycling
	// but not in the general case

	pathLengths := make([]int, 0, len(currentNodes))
	for _, node := range currentNodes {
		moves := 0
		for !strings.HasSuffix(node.Name, "Z") {
			node = processInstruction(node, instructions, moves)
			moves++
		}
		pathLengths = append(pathLengths, moves)
	}

	return strconv.Itoa(util.LowestCommonMultipleSlice(pathLengths))
}

// func allEndInZ(nodes []*Node) bool {
// 	for _, node := range nodes {
// 		if !strings.HasSuffix(node.Name, "Z") {
// 			return false
// 		}
// 	}
// 	return true
// }

type Node struct {
	Name     string
	Left     string
	Right    string
	Children [2]*Node // 0 == Left, 1==Right
}

func (current *Node) find(name string, visited []*Node) (*Node, bool) {
	if slices.Contains(visited, current) {
		return nil, false
	}
	visited = append(visited, current)

	if current.Name == name {
		return current, true
	}

	if current.Children[Left] != nil {
		if node, ok := current.Children[Left].find(name, visited); ok {
			return node, true
		}
	}

	if current.Children[Right] != nil {
		if node, ok := current.Children[Right].find(name, visited); ok {
			return node, true
		}
	}

	return nil, false
}

func (current *Node) findOrCreate(name string) *Node {
	if node, ok := current.find(name, make([]*Node, 0)); ok {
		return node
	} else {
		return &Node{Name: name, Children: [2]*Node{}}
	}
}

type Direction int

const (
	Left  = 0
	Right = 1
)

func parseInput(lines []string) ([]Direction, []*Node) {
	directions := parseDirections(strings.TrimSpace(lines[0]))
	nodes := parseNodeLines(lines[2:])

	return directions, nodes
}

func parseDirections(line string) (directions []Direction) {
	for _, s := range strings.Split(line, "") {
		switch s {
		case "R":
			directions = append(directions, Right)
		case "L":
			directions = append(directions, Left)
		default:
			panic("Not an R or L")
		}
	}
	return
}

func parseNodeLines(lines []string) []*Node {
	nodes := make([]*Node, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		under, left, right := parseNodeLine(line)
		nodes = append(nodes, &Node{Name: under, Left: left, Right: right, Children: [2]*Node{}})
	}
	slices.SortFunc(nodes, func(a, b *Node) int {
		return cmp.Compare(a.Name, b.Name)
	})

	for _, node := range nodes {
		n, found := slices.BinarySearchFunc(
			nodes,
			&Node{node.Left, "", "", [2]*Node{}},
			func(a *Node, b *Node) int {
				return cmp.Compare(a.Name, b.Name)
			},
		)
		if found {
			node.Children[Left] = nodes[n]
		}

		n, found = slices.BinarySearchFunc(
			nodes,
			&Node{node.Right, "", "", [2]*Node{}},
			func(a *Node, b *Node) int {
				return cmp.Compare(a.Name, b.Name)
			},
		)
		if found {
			node.Children[Right] = nodes[n]
		}
	}

	return nodes
}

func parseNodeLine(line string) (under string, left string, right string) {
	re := regexp.MustCompile(`^([A-Z0-9]{3}) = \(([A-Z0-9]{3}), ([A-Z0-9]{3})\)$`)
	matches := re.FindStringSubmatch(line)
	return matches[1], matches[2], matches[3]
}

func processInstruction(currentNode *Node, instructions []Direction, moves int) *Node {
	index := moves % len(instructions)
	instruction := instructions[index]
	return currentNode.Children[instruction]
}
