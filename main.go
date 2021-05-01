package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
)

type Rejection struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

var rejections []Rejection

func parse_csv() {
	file, err := os.Open("rejections.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for _, s := range lines {
		rejections = append(rejections, Rejection{Id: s[0], Message: s[1]})
	}
}

func getRejection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	random_message := rejections[rand.Intn(len(rejections))] // random index
	json.NewEncoder(w).Encode(random_message)                // JSON response

	log.Println("Response: ", random_message)
	return
}

func main() {
	parse_csv()

	r := mux.NewRouter()

	r.HandleFunc("/rejections", getRejection).Methods("GET")

	fmt.Println("Listening for requests at http://localhost:3000/rejections\n")
	log.Fatal(http.ListenAndServe(":3000", r))
}
