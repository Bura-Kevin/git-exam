package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", URLHandler)
	http.HandleFunc("/color", ColorHandler)

}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited: %s", r.URL.Path)
}

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}
