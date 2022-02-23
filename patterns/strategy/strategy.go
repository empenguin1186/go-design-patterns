package strategy

import (
	"encoding/json"
	"fmt"
)

type Searcher struct {
	searcherName string
	strategy     SearchStrategy
}

func NewSearcher(searcherName string, strategy SearchStrategy) *Searcher {
	return &Searcher{searcherName: searcherName, strategy: strategy}
}

func (c *Searcher) Execute(graph map[int][]int, start int, seen []bool) {
	result := c.strategy.Search(graph, start, seen)
	str, _ := json.Marshal(result)
	fmt.Printf("name: %s, result: %s\n", c.searcherName, string(str))
}

type SearchStrategy interface {
	Search(graph map[int][]int, start int, seen []bool) []int
}

type DepthFirstSearchStrategy struct{}

func NewDepthFirstSearchStrategy() *DepthFirstSearchStrategy {
	return &DepthFirstSearchStrategy{}
}

func (d *DepthFirstSearchStrategy) Search(graph map[int][]int, start int, seen []bool) []int {
	seen[start] = true
	var result []int
	result = append(result, start)

	list := graph[start]
	for _, e := range list {
		if seen[e] {
			continue
		}
		children := d.Search(graph, e, seen)
		result = append(result, children...)
	}
	return result
}

type BreadthFirstSearchStrategy struct {
	queue []int
}

func NewBreadthFirstSearchStrategy() *BreadthFirstSearchStrategy {
	return &BreadthFirstSearchStrategy{queue: make([]int, 0)}
}

func (b *BreadthFirstSearchStrategy) Search(graph map[int][]int, start int, seen []bool) []int {
	seen[start] = true
	b.queue = append(b.queue, start)

	var result []int
	for len(b.queue) != 0 {
		v := b.queue[0]
		b.queue = b.queue[1:]

		result = append(result, v)
		list := graph[v]
		for _, e := range list {
			if seen[e] {
				continue
			}
			b.queue = append(b.queue, e)
		}
	}
	return result
}
