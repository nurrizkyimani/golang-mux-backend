package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//this is comment to add
func HomeRoute(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintf(w, "You've requested the book: %s on", key)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/{key}", HomeRoute)
	http.ListenAndServe("127.0.0.1:8000", r)
}
