package dijkstra

import (
	"math"

	"github.com/emirpasic/gods/lists/arraylist"
)

func Dijkstra[T any](
	vertices []*T,
	source *T,
	neighborsOf func(node *T) []*T,
	edges func(a, b *T) float64,
	targetFound func(node *T) bool,
) (
	distances map[*T]float64,
	previousChain map[*T]*T,
) {
	distances = map[*T]float64{}
	previousChain = map[*T]*T{}

	byDistance := func(a, b interface{}) int {
		aAsserted := a.(*T)
		bAsserted := b.(*T)
		switch {
		case distances[aAsserted] > distances[bAsserted]:
			return 1
		case distances[aAsserted] < distances[bAsserted]:
			return -1
		default:
			return 0
		}
	}

	queue := arraylist.New()

	distances[source] = 0

	for _, node := range vertices {
		if node != source {
			distances[node] = math.Inf(1)
			previousChain[node] = nil
		}
		queue.Add(node)
	}

	for !queue.Empty() {
		queue.Sort(byDistance)

		value, _ := queue.Get(0)
		node := value.(*T)
		queue.Remove(0)

		if targetFound != nil {
			if targetFound(node) {
				return
			}
		}

		neighbors := neighborsOf(node)
		for _, neighbor := range neighbors {
			if !queue.Contains(neighbor) {
				continue
			}

			distance := 1.0
			if edges != nil {
				distance = edges(node, neighbor)
			}

			alt := distances[node] + distance
			if alt < distances[neighbor] {
				distances[neighbor] = alt
				previousChain[neighbor] = node
			}
		}
	}
	return
}
