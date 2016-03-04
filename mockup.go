package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var mockupResponse string

func init() {
	fileContent, _ := ioutil.ReadFile("mockup-response.json")
	mockupResponse = string(fileContent)
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	fmt.Fprintf(w, mockupResponse)
}

func main() {
	fmt.Println("Listening on localhost:8080. Ctrl+C to exit")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
