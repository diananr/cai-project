package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)


var mainDataset [][]float64

func makePrediction(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(res, "Error reading request body",
				http.StatusInternalServerError)
		}
		bodyJson := string(body)
		fmt.Println("bodyJson =>", bodyJson)

		rowToTest :=  mainDataset[3]
		predictionResponse := predictClassification(mainDataset, rowToTest, 5)

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
