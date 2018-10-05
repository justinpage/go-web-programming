package main

import (
	"html/template"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("X-XSS-Protection", "0") // disable protection in chrome
	t, _ := template.ParseFiles("8.tmpl.html")

	// e.g. <script>alert('Pwnd!');</script>
	t.Execute(w, template.HTML(r.FormValue("comment"))) // parse script

}

func form(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("8.form-for-submitting-xss-attack.html")
	t.Execute(w, nil)
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/process", process)
	http.HandleFunc("/form", form)
	server.ListenAndServe()
}
