package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/webhook", handleWebhook)
	log.Fatal(http.ListenAndServe(":80", nil))
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello Webhook!")
}
