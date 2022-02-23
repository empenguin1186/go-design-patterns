package main

import (
	"github.com/empenguin1186/go-design-patterns/patterns/strategy"
)

func main() {
	graph := map[int][]int{
		0: []int{1, 2},
		1: []int{3, 4},
		2: []int{5, 6},
	}

	dfsSearcher := strategy.NewSearcher("Depth First Search", strategy.NewDepthFirstSearchStrategy())
	bfsSearcher := strategy.NewSearcher("Breadth First Search", strategy.NewBreadthFirstSearchStrategy())

	dfsSearcher.Execute(graph, 0, make([]bool, 7))
	bfsSearcher.Execute(graph, 0, make([]bool, 7))
}
