package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strings"
)

type Rejection struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}

var rejections []Rejection

//go:embed rejections.csv
var content embed.FS

func parse_csv() {
	data, _ := content.ReadFile("rejections.csv") // read from embedded file
	stringdata := string(data)

	lines := strings.Split(stringdata, "\n")

	for _, line := range lines {
		split := strings.SplitN(line, ",", 2)              // SplitN limits substrings to 2, so it only splits on first comma in this case.
		split[1] = strings.Replace(split[1], "\"", "", -1) // replace double quotes in the string with nothing, if the final int input is < 0 then it replaces an infinite amount of quotes.
		rejections = append(rejections, Rejection{Id: split[0], Message: split[1]})
	}
}

func getRejection(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")       // JSON header so it displays properly in browsers
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
