package handler

import (
	"golangweb/entity"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	data := []entity.Product{
		{ID: 1, Name: "Stratos", Price: 12000000, Stock: 17},
		{ID: 2, Name: "Siskiu", Price: 8000000, Stock: 7},
		{ID: 3, Name: "Clovis", Price: 5000000, Stock: 1},
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}
}

func SabiliHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, my name is Sabili."))
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world! saya sedang belajar Golang."))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil || idNumber < 1 {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumber,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening", http.StatusInternalServerError)
		return
	}

	// fmt.Fprintf(w, "Product page : %d", idNumber)
}
