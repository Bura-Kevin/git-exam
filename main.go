package main

import (
	"fmt"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome - Index</h1>")
}

func main() {
	http.HandleFunc("/", IndexHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
