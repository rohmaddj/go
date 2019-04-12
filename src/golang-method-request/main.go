package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			w.Write([]byte("Here is post message"))
		} else if r.Method == "GET" {
			w.Write([]byte("Here is get message"))
		} else {
			http.Error(w, "", http.StatusBadRequest)
		}
	})

	fmt.Println("server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
