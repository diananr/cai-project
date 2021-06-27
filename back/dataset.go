package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
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

func getDataset() (dataset [][]string){
	urlDataSet := "https://github.com/diananr/cai-project/raw/master/dataset/casos_cai_2019.csv"

	data, err := readCSV(urlDataSet)
	if err != nil {
		fmt.Println("Error => ", err.Error())
		panic(err)
	}

	datasetDataCustom := [][]string{}
	for idx, row := range data {
		if idx == 0 {
			continue
		}

		rowFloat := strings.Split(row[0], ",")
		elementRow := []string{}

		for _, element := range rowFloat {
			elementRow = append(elementRow, element)
		}

		fmt.Println("row =>=>=>=>", elementRow)
		datasetDataCustom = append(datasetDataCustom, elementRow)
	}

	return datasetDataCustom;
}

func main() {
	getDataset()
}
