package dijkstra

import (
	"slices"
	"testing"
)

type element struct {
	name string
}

func buildPathFromPrevious(source *element, node *element, previousChain map[*element]*element) (result []*element) {
	if previousChain[node] != nil || node == source {
		for node != nil {
			result = append(result, node)
			node = previousChain[node]
		}
		slices.Reverse(result)
	}
	return
}
func stringifyPath(path []*element) (result string) {
	for _, node := range path {
		result += node.name
	}
	return
}

func TestDijkstraWithTarget(t *testing.T) {
	a := &element{"A"}
	b := &element{"B"}
	c := &element{"C"}
	d := &element{"D"}
	e := &element{"E"}
	f := &element{"F"}
	vertices := []*element{a, b, c, d, e, f}
	source := a
	neighborsOf := func(node *element) []*element {
		switch node.name {
		case "A":
			return []*element{b, c}
		case "B":
			return []*element{a, d}
		case "C":
			return []*element{a, f}
		case "D":
			return []*element{f, b, e}
		case "E":
			return []*element{d}
		case "F":
			return []*element{c, d}
		default:
			return []*element{}
		}
	}
	edgeDistanceBetween := func(nodeA *element, nodeB *element) float64 {
		return 1.0
	}
	targetFound := func(node *element) bool {
		return node == e
	}

	distances, previousChain := Dijkstra(vertices, source, neighborsOf, edgeDistanceBetween, targetFound)

	expectedDistanceToE := 3.0
	if expectedDistanceToE != distances[e] {
		t.Errorf("Expected: %f, Got %f", expectedDistanceToE, distances[e])
	}

	expectedPathToE := "ABDE"
	actualPathToE := stringifyPath(buildPathFromPrevious(a, e, previousChain))
	if expectedPathToE != actualPathToE {
		t.Errorf("Expected: %f, Got %f", expectedDistanceToE, distances[e])
	}

}

func TestDijkstraWithoutTarget(t *testing.T) {
	a := &element{"A"}
	b := &element{"B"}
	c := &element{"C"}
	d := &element{"D"}
	e := &element{"E"}
	f := &element{"F"}
	vertices := []*element{a, b, c, d, e, f}
	source := a
	neighborsOf := func(node *element) []*element {
		switch node.name {
		case "A":
			return []*element{b, c}
		case "B":
			return []*element{a, d}
		case "C":
			return []*element{a, f}
		case "D":
			return []*element{f, b, e}
		case "E":
			return []*element{d}
		case "F":
			return []*element{c, d}
		default:
			return []*element{}
		}
	}

	distances, previousChain := Dijkstra(vertices, source, neighborsOf, nil, nil)

	expectedDistancesFromA := map[*element]float64{
		a: 0.0,
		b: 1.0,
		c: 1.0,
		d: 2.0,
		e: 3.0,
		f: 2.0,
	}
	for node, distance := range expectedDistancesFromA {
		if distances[node] != distance {
			t.Errorf("Expected: %f, Got %f, for node %s", distance, distances[node], node.name)
		}
	}

	expectedPaths := map[*element]string{
		a: "A",
		b: "AB",
		c: "AC",
		d: "ABD",
		e: "ABDE",
		f: "ACF",
	}
	actualPaths := map[*element]string{
		a: stringifyPath(buildPathFromPrevious(a, a, previousChain)),
		b: stringifyPath(buildPathFromPrevious(a, b, previousChain)),
		c: stringifyPath(buildPathFromPrevious(a, c, previousChain)),
		d: stringifyPath(buildPathFromPrevious(a, d, previousChain)),
		e: stringifyPath(buildPathFromPrevious(a, e, previousChain)),
		f: stringifyPath(buildPathFromPrevious(a, f, previousChain)),
	}
	for node, path := range expectedPaths {
		if actualPaths[node] != path {
			t.Errorf("Expected: %s, Got %s for node %s", path, actualPaths[node], node.name)
		}
	}
}
