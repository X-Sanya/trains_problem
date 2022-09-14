package graph

import (
	"fmt"

	"github.com/x-sanya/trains_problem/train"
)

type Graph map[int][]*train.Train

func NewGraph(trains []train.Train) *Graph {
	g := make(Graph)
	for _, train := range trains {
		g.AddTrain(train.From, &train)
	}
	return &g
}

func (g *Graph) GetNodesAmount() int {
	return len(*g)
}

func (g *Graph) AddTrain(node int, t *train.Train) {
	_, ok := (*g)[node]
	if !ok {
		(*g)[node] = make([]*train.Train, 0)
	}
	trains := (*g)[node]
	newTrain := new(train.Train)
	*newTrain = *t
	(*g)[node] = append(trains, newTrain)
}

func (g *Graph) String() string {
	result := ""
	for numb, trains := range *g {
		result += fmt.Sprint(numb, len(trains), ":\n")
		for _, t := range trains {
			result += fmt.Sprintln("\t", t)
		}
	}
	return result
}
