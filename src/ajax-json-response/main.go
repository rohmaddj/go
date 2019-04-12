package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", ActionIndex)

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
	data := []struct {
		Name string
		Age  int
	}{
		{"Richard Grayson", 24},
		{"Jason Todd", 23},
		{"Tim Drake", 22},
		{"Damian Wayne", 21},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonInBytes)
}
