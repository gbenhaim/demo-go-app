package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there %s", r.URL.Path[1:])
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	port := ":8080"
	log.Printf("Listening on %s", port)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
