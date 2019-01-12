package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleDefaultPath)

	port := os.Getenv("PORT")

	log.Fatal(http.ListenAndServe(port, nil))
}

func handleDefaultPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love")
}
