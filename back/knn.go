package main

import (
	"fmt"
	"math"
	"sort"
)

type DistanceStruct struct {
	row      []float64
	distance   float64
}

func euclideanDistance(row1 []float64, row2 []float64) (distance float64){
	distance = 0.0
	for i := 0; i < len(row1)-1; i++ {
		distance += math.Pow(row1[i] - row2[i], 2)
	}
	return math.Sqrt(distance)
}

func getNeighbors(train [][]float64, testRow []float64, numNeighbors int) (neighbors [][]float64){
	distances := make([]DistanceStruct,0)

	for _, trainRow := range train{
		dist := euclideanDistance(testRow, trainRow)
		distStruct := DistanceStruct{row: trainRow,distance: dist}
		distances = append(distances, distStruct)
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	nearestNeighbors := [][]float64{}
	for i := 0; i < numNeighbors; i++ {
		nearestNeighbors = append(nearestNeighbors, distances[i].row)
	}
	return nearestNeighbors
}

func getMode(numbers []float64) (mode float64) {
		countMap := make(map[float64]int)
		for _, value := range numbers {
			countMap[value] += 1
		}

		max := 0
		for _, key := range numbers {
			freq := countMap[key]
			if freq > max {
				mode = key
				max = freq
			}
		}
		return;
	}

func predictClassification(train [][]float64, testRow []float64, numNeighbors int) (p float64){
	neighbors := getNeighbors(train, testRow, numNeighbors)
	output := []float64{}
	for _, row := range neighbors {
		output = append(output, row[len(row)-1])
	}
	prediction := getMode(output)
	return prediction;
}

func main() {
	dataset := [][] float64 {{2.7810836,2.550537003,0},
		{1.465489372,2.362125076,0},
		{3.396561688,4.400293529,0},
		{1.38807019,1.850220317,0},
		{3.06407232,3.005305973,0},
		{7.627531214,2.759262235,1},
		{5.332441248,2.088626775,1},
		{6.922596716,1.77106367,1},
		{8.675418651,-0.242068655,1},
		{7.673756466,3.508563011,1}}

	prediction := predictClassification(dataset, dataset[4], 3)
	fmt.Println("prediction: ", prediction)
}