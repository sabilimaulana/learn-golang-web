package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	aboutHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("This is about page."))
	}

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/sabili", sabiliHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to profile page."))
	})

	log.Print("Starting web on http://localhost:8080")

	err := http.ListenAndServe(":8080", mux)

	log.Fatal(err)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Welcome to Home Route."))
}

func sabiliHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Sabili."))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world! saya sedang belajar Golang."))
}
