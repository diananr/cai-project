package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Row struct {
	CAI   		float64  `json:"cai"`
	Age   		float64	 `json:"age"`
	Work   		float64  `json:"work"`
	Link   		float64  `json:"link"`
	Alcohol   float64  `json:"alcohol"`
	Smoke   	float64  `json:"smoke"`
	Drugs   	float64  `json:"drugs"`
	Addiction float64  `json:"addiction"`
	Risk   		float64  `json:"risk"`
	Month  	 	float64  `json:"month"`
}


var mainDataset [][]float64

func makePrediction(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Error reading request body",
				http.StatusInternalServerError)
		}

		var rowBody Row
		err = json.Unmarshal(body, &rowBody)
		if err != nil {
			http.Error(res, err.Error(), 500)
			return
		}

		rowToTest :=  mainDataset[3]
		predictionResponse := predictClassification(mainDataset, rowToTest, 5)
		fmt.Println("predictionResponse =>", predictionResponse)

	} else {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
	}

	res.Header().Set("Content-Type", "application/json")

}

func handleRequest() {
	http.HandleFunc("/prediction", makePrediction)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	mainDataset = getDataset()
	handleRequest()
}
