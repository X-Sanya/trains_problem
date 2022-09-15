package main

import (
	"fmt"
	"time"

	"github.com/x-sanya/trains_problem/ants_algorithm"
	"github.com/x-sanya/trains_problem/graph"
)

func main() {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println("Something was wrong:", str)
		}
	}()

	var dataFileName = "test_task_data.xlsx"
	var dataSheetName = "test_task_data"

	trains, err := UploadData(dataFileName, dataSheetName)

	if err != nil {
		panic(err)
	}

	g := graph.NewGraph(trains)
	a := ants_algorithm.NewAntsSwarm(trains, g, len(*g), 2.0, 1.0, 0.3, float64(time.Hour*20))
	shortestRoute, shortestTime := a.FindShortestRoute(1000)

	fmt.Printf("The shortest path is found with time %v: \n", shortestTime)
	for _, t := range shortestRoute {
		fmt.Println(&trains[t])
	}

	a.ChangeSetting(len(*g), 1.0, 1.5, 0.3, 100)
	cheapestRoute, lowestPrice := a.FindCheapestRoute(10000)

	fmt.Printf("The cheapest path is found with price %.2f: \n", lowestPrice)
	for _, t := range cheapestRoute {
		fmt.Println(&trains[t])
	}
	var wait byte
	fmt.Scan(&wait)
}
