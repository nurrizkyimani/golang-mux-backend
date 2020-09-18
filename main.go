package main

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

// this is comment to add
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintf(w, "You've requested the book: %s on", key)
}

type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]

	for _, article := range Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

var Articles []Article

func main() {

	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "3", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	r := mux.NewRouter()
	r.HandleFunc("/{key}", HomeHandler)
	r.HandleFunc("/product/{key}", ProductHandler)
	// r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
