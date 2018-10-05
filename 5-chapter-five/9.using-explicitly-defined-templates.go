package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("9.layout.html")
	t.ExecuteTemplate(w, "layout", "")
}

func random(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("9.layout.html", "9.red_hello.html")
	} else {
		t, _ = template.ParseFiles("9.layout.html", "9.blue_hello.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func block(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles("9.layout-block.html", "9.red_hello.html")
	} else {
		t, _ = template.ParseFiles("9.layout-block.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/random", random)
	http.HandleFunc("/block", block)
	server.ListenAndServe()
}
