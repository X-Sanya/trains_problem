package main

import (
	"fmt"

	"github.com/x-sanya/trains_problem/ants_algorithm"
	"github.com/x-sanya/trains_problem/graph"
)

func main() {
	var dataFileName = "test_task_data.xlsx"
	var dataSheetName = "test_task_data"

	trains, err := UploadData(dataFileName, dataSheetName)

	fmt.Println(err)

	g := graph.NewGraph(trains)
	a := ants_algorithm.NewAntsSwarm(1, 1, 30, 6, 0.3)
	a.FindBestRoute(g)
	fmt.Println(g)
}
