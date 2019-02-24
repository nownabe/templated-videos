package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", newHomeHandler())
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/render", renderHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "posted")
}

func renderHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "render")
}
