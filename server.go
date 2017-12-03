package main

import (
	"context"
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("Starting server...")
	server := &http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		log.Printf("Request recieved from '%q'", r.URL.Path)
	})
	http.HandleFunc("/close", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Closing server...")
		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		server.Shutdown(ctx)
	})
	log.Fatal(server.ListenAndServe())
	log.Println("Server exited succesfully")
}
