package main

import (
	"github.com/rebay1982/gostack/handler"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	// Prevents the fixed path / from acting as a catch all.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return // If we didn't return here, it would write the message at the bottom too.
	}

	w.Write([]byte("Hello from Gostack"))
}

func foo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, FOO"))
}

func post(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {

		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", 405)

		return
	}

	w.Write([]byte("POST Accepted"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)     // Fix path pattern
	mux.HandleFunc("/foo/", foo)  // Subtree pattern
	mux.HandleFunc("/post", post) // Fixed path pattern.
	mux.HandleFunc("/users", handler.Users)

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
