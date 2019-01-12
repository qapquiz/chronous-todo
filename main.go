package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", handleDefaultPath)

	port, err := os.Getenv("PORT")
	if err != nil {

	}
	log.Fatal(http.ListenAndServe(port, nil))
}

func handleDefaultPath(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love")
}
