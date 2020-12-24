package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

//HomeHandler is a function
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to home page"))
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
