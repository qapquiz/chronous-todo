package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleDefaultPath)
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func handleDefaultPath(w http.ResponseWriter, r *http.Request) {
	log.Println(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Webhook!")
}
