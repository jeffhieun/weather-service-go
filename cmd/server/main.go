package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", weatherHandler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"message": "Welcome to the Weather Service! Use /weather?location=YOUR_CITY to get weather information."}`)
}
