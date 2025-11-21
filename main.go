package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", URLHandler)

}

func URLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You visited: %s", r.URL.Path)
}
