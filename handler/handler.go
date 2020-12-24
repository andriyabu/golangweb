package handler

import (
	"golangweb/entity"
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

	tmpl, err := template.ParseFiles(path.Join("views", "index.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "Error is happening, keep calm and drink a coffee", http.StatusInternalServerError)
		return
	}

	// data := map[string]interface{}{
	// 	"title":   "Golang web basic",
	// 	"content": "Learning Golang web from buildwithangga.com",
	// }

	// data := entity.Product{ID: 1, Name: "Mobilio", Price: 200000000, Stock: 5}
	data := []entity.Product{
		{ID: 1, Name: "Mobilio", Price: 200000000, Stock: 5},
		{ID: 1, Name: "Expander", Price: 280000000, Stock: 10},
		{ID: 1, Name: "Tucson", Price: 180000000, Stock: 2},
		{ID: 1, Name: "Escape", Price: 1600000000, Stock: 4},
		{ID: 1, Name: "Fortuner", Price: 500000000, Stock: 8},
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

	// fmt.Fprintf(w, "Product page : %d", idNumb)
	tmpl, err := template.ParseFiles(path.Join("views", "product.html"), path.Join("views", "layout.html"))

	if err != nil {
		log.Println(err)
		http.Error(w, "keep calm and drink coffee", http.StatusInternalServerError)
		return
	}

	data := map[string]interface{}{
		"content": idNumb,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		log.Println(err)
		http.Error(w, "keep calm and drink coffee", http.StatusInternalServerError)
		return
	}
}
