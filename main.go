package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Dit is een Go webserver. Hallo vanuit Go!")
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})
	mux.HandleFunc("/time", func(w http.ResponseWriter, r *http.Request) {
		currentTime := time.Now().Format(time.RFC1123)
		fmt.Fprintf(w, "Huidige server tijd: %s\n", currentTime)
	})

	// Lees poort uit omgevingsvariabele, standaard naar 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Luisteren op poort %s in de container…", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
