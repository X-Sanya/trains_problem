package ants_algorithm

import (
	"fmt"
	"time"

	"github.com/x-sanya/trains_problem/graph"
	"github.com/x-sanya/trains_problem/train"
)

type AntsSwarm struct {
	a, b, q     int
	antsAmount  int
	evaporation float32
}

func NewAntsSwarm(a, b, q, antsAmount int, evaporation float32) *AntsSwarm {
	return &AntsSwarm{a: a, b: b, q: q, antsAmount: antsAmount, evaporation: evaporation}
}

func (a *AntsSwarm) FindBestRoute(g *graph.Graph) ([]*train.Train, time.Duration) {
	bestRoute := make([]*train.Train, 0, g.GetNodesAmount())
	trains := a.pheromoneInitialization(g)
	_, ok := trains[nil]
	fmt.Println(ok)
	//ants run
	//change pheromones
	return bestRoute, time.Duration(0)
}

func (a *AntsSwarm) pheromoneInitialization(g *graph.Graph) map[*train.Train]float32 {
	trainsP := make(map[*train.Train]float32)
	var ph float32 = 0.1

	for _, trains := range *g {
		for _, t := range trains {
			trainsP[t] = ph
		}
	}
	return trainsP
}
