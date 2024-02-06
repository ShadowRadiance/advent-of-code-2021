package day23

import (
	"slices"
	"strconv"
	"strings"

	"github.com/emirpasic/gods/queues/arrayqueue"

	"github.com/shadowradiance/advent-of-code/2023-go/util/grids"
)

type Solution struct{}

type gridT = grids.Grid[rune]
type posT = grids.Vector2D[int]
type dirT = grids.Vector2D[int]
type tileT = rune

const (
	Wall    tileT = '#'
	Path    tileT = '.'
	SlopeLt tileT = '<'
	SlopeRt tileT = '>'
	SlopeUp tileT = '^'
	SlopeDn tileT = 'v'
)

type graphNodeT struct {
	pos  posT
	tile tileT
	// edges []*graphNodeT
	reducedEdges map[*graphNodeT]int
}

func (Solution) Part01(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := initializeGrid(lines)
	graph := initializeGraph(
		firstPathPosIn(grid, 0),
		grid,
		make(map[posT]*graphNodeT))

	// fmt.Println(grid.Dump())
	longestPath := determineMaxDepth(graph, make([]*graphNodeT, 0), 0, 0)

	return strconv.Itoa(longestPath)
}

func (Solution) Part02(input string) string {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	if len(lines) == 0 {
		return "NO DATA"
	}

	grid := initializeGrid(lines)
	removeSlopes(grid)
	graph := initializeGraph(
		firstPathPosIn(grid, 0),
		grid,
		make(map[posT]*graphNodeT),
	)
	reduceEdges(graph)

	// fmt.Println(grid.Dump())
	longestPath := longestPathWithoutLoops(graph, grid)

	return strconv.Itoa(longestPath)
}

func initializeGrid(lines []string) *gridT {
	grid := grids.NewGrid(lines)
	return &grid
}

func initializeGraph(root posT, grid *gridT, existingNodes map[posT]*graphNodeT) *graphNodeT {
	tile := grid.AtPos(root)
	if tile == Wall {
		return nil
	}
	node := &graphNodeT{
		pos:  root,
		tile: tile,
		// edges: make([]*graphNodeT, 0),
		reducedEdges: make(map[*graphNodeT]int),
	}
	existingNodes[root] = node
	for _, neighborPos := range neighborsOfPos(root, grid) {
		if neighborNode, ok := existingNodes[neighborPos]; ok {
			node.reducedEdges[neighborNode] = 1
			// node.edges = append(node.edges, neighborNode)
			continue
		}
		neighborNode := initializeGraph(neighborPos, grid, existingNodes)
		if neighborNode != nil {
			node.reducedEdges[neighborNode] = 1
			// node.edges = append(node.edges, neighborNode)
		}
	}
	return node
}

func firstPathPosIn(grid *gridT, y int) posT {
	row := grid.RowAt(y)
	for x, char := range row {
		tile := char
		if tile == Path {
			return posT{X: x, Y: y}
		}
	}
	panic("No path tile in row" + strconv.Itoa(y))
}

func directions() []dirT {
	return []dirT{
		grids.East[int](),
		grids.South[int](),
		grids.West[int](),
		grids.North[int](),
	}
}

func neighborsOfPos(pos posT, grid *gridT) (positions []posT) {
	for _, dir := range directions() {
		tile := grid.AtPos(pos)
		if (tile == SlopeDn && dir != grids.South[int]()) ||
			(tile == SlopeUp && dir != grids.North[int]()) ||
			(tile == SlopeLt && dir != grids.West[int]()) ||
			(tile == SlopeRt && dir != grids.East[int]()) {
			continue
		}
		test := pos.Add(dir)
		if test.InBounds(0, 0, grid.Width()-1, grid.Height()-1) {
			// exclude walking into walls
			if grid.AtPos(test) == Wall {
				continue
			}
			positions = append(positions, test)
		}
	}
	return
}

func determineMaxDepth(graph *graphNodeT, visited []*graphNodeT, depth int, maxSoFar int) int {
	maxDepth := max(depth, maxSoFar)

	visited = append(visited, graph)
	for edge := range graph.reducedEdges {
		// for _, edge := range graph.edges {
		if slices.Contains(visited, edge) {
			continue
		}
		maxDepth = max(maxDepth, determineMaxDepth(edge, visited, depth+1, maxDepth))
	}

	return maxDepth
}

func removeSlopes(grid *gridT) {
	for y, row := range *grid {
		for x, cell := range row {
			switch tileT(cell) {
			case Wall:
				continue
			default:
				grid.SetAt(x, y, Path)
			}
		}
	}
}

type visitedT = grids.Grid[bool]

func dfs(node *graphNodeT, visited *visitedT, pos posT, dst posT, steps int, maxSteps *int) {

	if pos == dst {
		*maxSteps = max(*maxSteps, steps)
	}
	if visited.AtPos(pos) {
		return
	}

	visited.SetAtPos(pos, true)
	for edgeNode, enSteps := range node.reducedEdges {
		dfs(edgeNode, visited, edgeNode.pos, dst, steps+enSteps, maxSteps)
	}
	// for _, edgeNode := range node.edges {
	// 	dfs(edgeNode, visited, edgeNode.pos, dst, steps+1, maxSteps)
	// }
	visited.SetAtPos(pos, false)

}

func longestPathWithoutLoops(graph *graphNodeT, grid *gridT) int {
	maxSteps := 0

	visited := grids.New[bool](grid.Height(), grid.Width())

	dfs(
		graph,
		&visited,
		firstPathPosIn(grid, 0),
		firstPathPosIn(grid, grid.Height()-1),
		0,
		&maxSteps,
	)

	return maxSteps
}

func reduceEdges(graph *graphNodeT) {
	visited := make([]*graphNodeT, 0)

	q := arrayqueue.New()
	q.Enqueue(graph)

	for !q.Empty() {
		value, _ := q.Dequeue()
		node, _ := value.(*graphNodeT)
		if slices.Contains(visited, node) {
			continue
		}
		visited = append(visited, node)
		for gn := range node.reducedEdges {
			q.Enqueue(gn)
		}

		if len(node.reducedEdges) == 2 {
			// if there is exactly one "in" and one "out"
			// a <-n-> graph <-m-> b
			// a <-n+m-> b

			var aGraph, bGraph *graphNodeT
			var aDist, bDist int
			for g, dist := range node.reducedEdges {
				if aGraph == nil {
					aGraph = g
					aDist = dist
				} else {
					bGraph = g
					bDist = dist
				}
			}
			delete(aGraph.reducedEdges, node)
			delete(bGraph.reducedEdges, node)
			aGraph.reducedEdges[bGraph] = aDist + bDist
			bGraph.reducedEdges[aGraph] = aDist + bDist
		}
	}

}
