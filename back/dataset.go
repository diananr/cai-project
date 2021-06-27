package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func readCSV(url string) ([][]string, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error getting URL => ", err.Error())
		return nil, err
	}

	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error while reading the file => ", err.Error())
		return nil, err
	}

	return data, nil
}

func getDataset() (dataset [][]float64){
	urlDataSet := "https://github.com/diananr/cai-project/raw/master/dataset/casos_cai_2019.csv"

	data, err := readCSV(urlDataSet)
	if err != nil {
		fmt.Println("Error => ", err.Error())
		panic(err)
	}

	datasetDataCustom := [][]float64{}
	for idx, row := range data {
		if idx == 0 {
			continue
		}

		rowFloat := strings.Split(row[0], ",")
		elementRow := []float64{}

		for _, element := range rowFloat {
			valueFloat, _ := strconv.ParseFloat(element, 64)
			elementRow = append(elementRow, valueFloat)
		}

		datasetDataCustom = append(datasetDataCustom, elementRow)
	}

	return datasetDataCustom;
}
