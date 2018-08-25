package main

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("parse", request.ParseForm())
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", Handler)
	http.ListenAndServe(":8080", nil)
}
