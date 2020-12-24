package handler

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path"
	"strconv"
)

//HomeHandler is a function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	// w.Write([]byte("Welcome to home page"))

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm and drink a coffee", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"title":   "Golang web basic",
		"content": "Learning Golang web from buildwithangga.com",
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm and drink a coffee", http.StatusInternalServerError)
		return
	}
}

//HelloHandler is a function
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heloo saya lagi belaja web golang"))

}

//AboutHandler is a function
func AboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='color: red'> About Page </h1>"))
}

//ProductHandler is a function
func ProductHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product page : %d", idNumb)
}
