package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("1.a-simple-template.html")
	t.Execute(w, "Hello World!")
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
