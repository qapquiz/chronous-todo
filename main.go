package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleDefaultPath)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleDefaultPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love")
}
