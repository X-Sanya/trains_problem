package ants_algorithm

import (
	"math"
	"math/rand"
	"sync"
	"time"

	"github.com/x-sanya/trains_problem/graph"
	"github.com/x-sanya/trains_problem/train"
)

type AntsSwarm struct {
	m                    sync.Mutex
	w                    sync.WaitGroup
	trains               []train.Train
	g                    *graph.Graph
	evaporation, q, a, b float64
	antsAmount           int
}

func NewAntsSwarm(trains []train.Train, g *graph.Graph, antsAmount int, a, b, evaporation, q float64) *AntsSwarm {
	return &AntsSwarm{trains: trains, g: g, a: a, b: b, q: q, antsAmount: antsAmount, evaporation: evaporation}
}

func (antsSwarm *AntsSwarm) ChangeSetting(antsAmount int, a, b, evaporation, q float64) {
	antsSwarm.a, antsSwarm.b, antsSwarm.evaporation, antsSwarm.q, antsSwarm.antsAmount = a, b, evaporation, q, antsAmount
}

func (a *AntsSwarm) FindShortestRoute(repeat int) ([]int, time.Duration) {
	rand.Seed(int64(time.Now().Nanosecond()))
	var bestRoute []int
	bestTime := time.Duration(time.Hour * 24 * 20)
	pheromones := a.pheromoneInitialization(0.1)

	for i := 0; i < repeat; i++ {
		for ant := 0; ant < a.antsAmount; ant++ {
			a.w.Add(1)
			go a.runTime(pheromones, &bestRoute, &bestTime)
		}
		a.w.Wait()
		for _, ph := range pheromones {
			ph[0] = (ph[0])*(1-a.evaporation) + ph[1]
			if ph[0] > 0.9999 {
				ph[0] = 0.9999
			}
			ph[1] = 0
		}
	}
	return bestRoute, bestTime
}

func (a *AntsSwarm) FindCheapestRoute(repeat int) ([]int, float32) {
	var bestRoute []int
	bestPrice := float32(math.MaxFloat32)
	pheromones := a.pheromoneInitialization(0.1)

	for i := 0; i < repeat; i++ {

		for ant := 0; ant < a.antsAmount; ant++ {
			a.w.Add(1)
			go a.runPrice(pheromones, &bestRoute, &bestPrice)
		}
		a.w.Wait()
		for _, ph := range pheromones {
			ph[0] = (ph[0])*(1-a.evaporation) + ph[1]
			if ph[0] > 0.9999 {
				ph[0] = 0.9999
			}
			ph[1] = 0
		}
	}
	return bestRoute, bestPrice
}

func (a *AntsSwarm) pheromoneInitialization(pheromone float64) [][2]float64 {
	pheromones := make([][2]float64, len(a.trains))
	for i := range a.trains {
		pheromones[i] = [2]float64{pheromone, 0}
	}
	return pheromones
}

func (a *AntsSwarm) runTime(pheromones [][2]float64, bestRoute *[]int, bestTime *time.Duration) {
	defer a.w.Done()
	currentTrainIndex := rand.Intn(len(a.trains))
	currentTrain := &a.trains[currentTrainIndex]
	currentStation := currentTrain.To
	availableStation := a.g.GetNodes()
	route := make([]int, 1, len(availableStation))
	route[0] = currentTrainIndex
	delete(availableStation, currentStation)
	delete(availableStation, currentTrain.From)
	currentTime := currentTrain.EndTime
	resultTime := getDurationBetween(currentTrain.StartTime, currentTime)
	for len(availableStation) != 0 {
		desires := make(map[int]float64, len((*a.g)[currentStation]))
		sumDesires := 0.0
		for _, t := range (*a.g)[currentStation] {
			nextTrain := a.trains[t]
			if !availableStation[nextTrain.To] {
				continue
			}
			dur := getDurationBetween(currentTime, nextTrain.StartTime) + getDurationBetween(nextTrain.StartTime, nextTrain.EndTime)
			desires[t] = math.Pow(float64(pheromones[t][0]), a.a) * math.Pow(float64(time.Hour*6)/float64(dur), a.b)
			sumDesires += desires[t]
		}
		if len(desires) == 0 {
			break
		}
		key := rand.Float64()
		currentTrainIndex = FindNearestT(desires, sumDesires, key)
		currentTrain = &a.trains[currentTrainIndex]
		dur := getDurationBetween(currentTime, currentTrain.StartTime) + getDurationBetween(currentTrain.StartTime, currentTrain.EndTime)
		resultTime += dur
		currentTime = currentTrain.EndTime
		currentStation = currentTrain.To
		route = append(route, currentTrainIndex)
		delete(availableStation, currentStation)
	}
	if len(availableStation) != 0 {
		return
	}
	if resultTime <= *bestTime {
		a.m.Lock()
		*bestTime = resultTime
		*bestRoute = route
		a.m.Unlock()
	}
	for _, t := range route {
		a.m.Lock()
		pheromones[t][1] += a.q / float64(resultTime)
		a.m.Unlock()
	}
}

func (a *AntsSwarm) runPrice(pheromones [][2]float64, bestRoute *[]int, bestPrice *float32) {
	defer a.w.Done()
	currentTrainIndex := rand.Intn(len(a.trains))
	currentTrain := &a.trains[currentTrainIndex]
	currentStation := currentTrain.To
	resultPrice := currentTrain.Price
	availableStation := a.g.GetNodes()
	delete(availableStation, currentStation)
	delete(availableStation, currentTrain.From)
	route := make([]int, 1, len(availableStation))
	route[0] = currentTrainIndex

	for len(availableStation) != 0 {
		desires := make(map[int]float64, len((*a.g)[currentStation]))
		sumDesires := 0.0
		for _, t := range (*a.g)[currentStation] {
			nextTrain := a.trains[t]
			if !availableStation[nextTrain.To] {
				continue
			}
			price := nextTrain.Price
			desires[t] = math.Pow(float64(pheromones[t][0]), a.a) * math.Pow(float64(50.0)/float64(price), a.b)
			sumDesires += desires[t]
		}
		if len(desires) == 0 {
			break
		}
		key := rand.Float64()
		currentTrainIndex := FindNearestT(desires, sumDesires, key)
		currentTrain = &a.trains[currentTrainIndex]
		resultPrice += currentTrain.Price
		currentStation = currentTrain.To
		route = append(route, currentTrainIndex)
		delete(availableStation, currentStation)
	}
	if len(availableStation) != 0 {
		return
	}
	if resultPrice <= *bestPrice {
		a.m.Lock()
		*bestPrice = resultPrice
		*bestRoute = route
		a.m.Unlock()
	}
	for _, t := range route {
		a.m.Lock()
		pheromones[t][1] += a.q / float64(resultPrice)
		a.m.Unlock()
	}
}

func FindNearestT(keys map[int]float64, den float64, key float64) int {
	T := 0
	min := 2.0
	for i, val := range keys {
		val /= den
		dif := math.Abs(val - key)
		if dif <= min {
			min = dif
			T = i
		}
	}
	return T
}

func getDurationBetween(start, end time.Time) time.Duration {
	hour, min, sec := start.Clock()
	dur := time.Hour*time.Duration(hour) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec)
	resultTime := end.Add(-dur)
	hour, min, sec = resultTime.Clock()
	dur = time.Hour*time.Duration(hour) + time.Minute*time.Duration(min) + time.Second*time.Duration(sec)
	return dur
}
