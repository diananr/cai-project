package main

import (
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
