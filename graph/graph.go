package graph

import (
	"github.com/x-sanya/trains_problem/train"
)

type Graph map[int][]int

func NewGraph(trains []train.Train) *Graph {
	g := make(Graph)
	for i, train := range trains {
		g.AddTrain(train.From, i)
	}
	return &g
}

func (g *Graph) GetNodes() map[int]bool {
	nodes := make(map[int]bool, len(*g))
	for node := range *g {
		nodes[node] = true
	}
	return nodes
}

func (g *Graph) AddTrain(node, t int) {
	_, ok := (*g)[node]
	if !ok {
		(*g)[node] = make([]int, 0)
	}
	trains := (*g)[node]
	(*g)[node] = append(trains, t)
}
