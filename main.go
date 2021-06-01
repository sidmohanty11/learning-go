package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/mux"

	// "math/rand"
	"net/http"
	"strconv"
)

//author struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

//Book struct (Model)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

var b []Book

func main() {
	//Init Router
	r := mux.NewRouter()

	//mock data
	for i := 0; i < 10; i++ {
		b = append(b, Book{
			ID:     strconv.Itoa(i),
			Isbn:   strconv.Itoa(i * 1000),
			Title:  "Influence",
			Author: &Author{Firstname: "Robert", Lastname: "Cialdini"},
		})
	}

	//Route handlers -> Endpts
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", addBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", r))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get params

	for _, item := range b {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func addBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}
