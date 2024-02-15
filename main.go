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

var banner string = `
██████╗ ███████╗     ██╗███████╗ ██████╗████████╗██╗ ██████╗ ███╗   ██╗
██╔══██╗██╔════╝     ██║██╔════╝██╔════╝╚══██╔══╝██║██╔═══██╗████╗  ██║
██████╔╝█████╗       ██║█████╗  ██║        ██║   ██║██║   ██║██╔██╗ ██║
██╔══██╗██╔══╝  ██   ██║██╔══╝  ██║        ██║   ██║██║   ██║██║╚██╗██║
██║  ██║███████╗╚█████╔╝███████╗╚██████╗   ██║   ██║╚██████╔╝██║ ╚████║
╚═╝  ╚═╝╚══════╝ ╚════╝ ╚══════╝ ╚═════╝   ╚═╝   ╚═╝ ╚═════╝ ╚═╝  ╚═══╝

 █████╗ ██████╗ ██╗
██╔══██╗██╔══██╗██║
███████║██████╔╝██║
██╔══██║██╔═══╝ ██║
██║  ██║██║     ██║
╚═╝  ╚═╝╚═╝     ╚═╝
`

//go:embed rejections.csv
var content embed.FS

func parse_csv() {
	data, _ := content.ReadFile("rejections.csv") // read from embedded file
	stringdata := string(data)

	lines := strings.Split(stringdata, "\n")

	for _, line := range lines {
		id, message, _ := strings.Cut(line, ",")         // split string on first instance of separator (comma).
		message = strings.Replace(message, "\"", "", -1) // replace double quotes in the string with nothing, if the final int input is < 0 then it replaces an infinite amount of quotes.
		rejections = append(rejections, Rejection{Id: id, Message: message})
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
	fmt.Println(banner)
	parse_csv()

	r := mux.NewRouter()

	r.HandleFunc("/rejections", getRejection).Methods("GET")

	fmt.Println("Listening for requests at http://localhost:3000/rejections\n")
	log.Fatal(http.ListenAndServe(":3000", r))
}
