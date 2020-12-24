package main

import (
	"golangweb/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// create passing a function to variable as a handler
	profileHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Profile Page"))
	}
	//routing
	mux.HandleFunc("/", handler.HomeHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/about", handler.AboutHandler)
	mux.HandleFunc("/profile", profileHandler)
	mux.HandleFunc("/product", handler.ProductHandler)
	//function in routing parameter
	mux.HandleFunc("/service", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Our Service"))
	})
	log.Println("Starting Web server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)

}
