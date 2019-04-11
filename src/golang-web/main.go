package main

import "fmt"
import "net/http"
import "html/template"
import "path"

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))
			
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		filepath := path.Join("views", "index.html")
		tmpl, err := template.ParseFiles(filepath)
		
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"title": "Learning Golang Web",
			"name": "Batman",
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server running at localhost:9000")
	http.ListenAndServe(":9000", nil)
}