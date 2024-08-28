package main

import (
	"fmt"
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
	fmt.Fprintf(w, "Snippet view with id: %d", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Snippet create %s", time.Now())
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "GO")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Save success"))
}
