package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	// create passing a function to variable as a handler
	profileHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	}
	//routing
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/product", productHandler)
	//function in routing parameter
	mux.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Our Service"))
	})
	log.Println("Starting Web server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Welcome to home page"))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Heloo saya lagi belaja web golang"))

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1 style='color: red'> About Page </h1>"))
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	idNumb, err := strconv.Atoi(id)

	if err != nil || idNumb < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Product page : %d", idNumb)
}
