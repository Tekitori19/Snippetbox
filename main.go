package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	// Atoi convert string to int
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Snippet view with id: %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Snippet create %s", time.Now())
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save success"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", greet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("Server start at http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
