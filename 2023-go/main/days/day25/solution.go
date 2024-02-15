package day25

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/queues/arrayqueue"

	"github.com/shadowradiance/advent-of-code/2023-go/util"
)

type StringSet = map[string]bool

type StringPair struct {
	first  string
	second string
}

type Node struct {
	name      string
	connected StringSet
}

type Edge struct {
	lft  string
	rgt  string
	oLft string
	oRgt string
}

type NodeMap map[string]*Node

func (n NodeMap) Clone() (clone NodeMap) {
	clone = NodeMap{}
	for k, node := range n {
		connected := StringSet{}
		for s, b := range node.connected {
			connected[s] = b
		}
		clone[k] = &Node{name: node.name, connected: connected}
	}
	return
}

func parseLines(lines []string) (nodes NodeMap, edges []*Edge) {
	inputData := map[string][]string{}
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		rootName := parts[0]
		connectedNames := strings.Split(strings.TrimSpace(parts[1]), " ")
		if _, ok := inputData[rootName]; !ok {
			inputData[rootName] = make([]string, 0)
		}
		inputData[rootName] = append(inputData[rootName], connectedNames...)
	}

	nodes = map[string]*Node{}
	for key, list := range inputData {
		nodes[key] = &Node{name: key, connected: StringSet{}}
		for _, name := range list {
			nodes[name] = &Node{name: name, connected: StringSet{}}
		}
	}

	for key, list := range inputData {
		left := nodes[key]
		for _, name := range list {
			right := nodes[name]
			left.connected[right.name] = true
			right.connected[left.name] = true
			edges = append(edges, &Edge{lft: left.name, rgt: right.name, oLft: left.name, oRgt: right.name})
		}
	}

	return
}

func tryCollapse(nodeMap NodeMap, edges []*Edge, want int) ([]StringPair, bool) {
	for len(edges) > want {
		randomIndex := rand.Intn(len(edges))
		rr := edges[randomIndex]

		edges = slices.Delete(edges, randomIndex, randomIndex+1)
		firstNode := nodeMap[rr.lft]
		secondNode := nodeMap[rr.rgt]

		delete(firstNode.connected, secondNode.name)
		delete(secondNode.connected, firstNode.name)

		indexes := make([]int, 0)
		for i, p := range edges {
			if (rr.lft == p.lft || rr.lft == p.rgt) && (rr.rgt == p.lft || rr.rgt == p.rgt) {
				indexes = append(indexes, i)
			}
		}
		slices.Sort(indexes)
		slices.Reverse(indexes)

		for _, index := range indexes {
			edges = slices.Delete(edges, index, index+1)
		}

		for s := range secondNode.connected {
			otherNode := nodeMap[s]
			delete(otherNode.connected, secondNode.name)
			otherNode.connected[firstNode.name] = true
			firstNode.connected[otherNode.name] = true
		}

		for _, pair := range edges {
			if pair.lft == secondNode.name {
				pair.lft = firstNode.name
			} else if pair.rgt == secondNode.name {
				pair.rgt = firstNode.name
			}
		}
	}

	if len(edges) != 3 {
		return make([]StringPair, 0), false
	} else {
		pp := util.Transform(
			edges,
			func(edge *Edge) StringPair { return StringPair{edge.oLft, edge.oRgt} },
		)
		return pp, true
	}
}

func collapse(nodeMap NodeMap, edges []*Edge, want int, maxTries int) (found []StringPair, ok bool) {
	for iteration := 0; iteration < maxTries; iteration++ {
		nodesClone := nodeMap.Clone()
		edgesClone := make([]*Edge, len(edges))
		for i, edge := range edges {
			edgesClone[i] = &Edge{lft: edge.lft, rgt: edge.rgt, oLft: edge.oLft, oRgt: edge.oRgt}
		}

		if found, ok = tryCollapse(nodesClone, edgesClone, want); ok {
			fmt.Println("found after:", iteration, "iterations")
			return
		}
	}
	return make([]StringPair, 0), false
}

type Solution struct{}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	nodes, edges := parseLines(lines)
	found, ok := collapse(nodes, edges, 3, 50_000)
	if !ok {
		fmt.Println("Crap")
		return "FAILED"
	}

	// remove the three nodes
	for _, pair := range found {
		left := nodes[pair.first]
		right := nodes[pair.second]
		delete(left.connected, right.name)
		delete(right.connected, left.name)
	}

	// grab a node on the left-hand side
	leftSet := StringSet{}

	root := nodes[found[0].first]
	q := arrayqueue.New()
	q.Enqueue(root.name)
	for !q.Empty() {
		item, _ := q.Dequeue()
		elem := item.(string)
		if !leftSet[elem] {
			leftSet[elem] = true
			node := nodes[elem]
			for name := range node.connected {
				q.Enqueue(name)
			}
		}
	}

	rightSet := StringSet{}
	for s := range nodes {
		if !leftSet[s] {
			rightSet[s] = true
		}
	}

	result := len(leftSet) * len(rightSet)

	return strconv.Itoa(result)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	return "PENDING"
}
