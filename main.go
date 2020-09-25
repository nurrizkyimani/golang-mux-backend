package main

import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"

	"github.com/gorilla/mux"
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"github.com/nurrizkyimani/golang-mux-backend/service"
	// "github.com/nurrizkyimani/golang-mux-backend/model"

	// "gorm.io/gorm"s

	// "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// "github.com/joho/godotenv"
)





func main() {
	//Initialize datbase
	db.InitDatabase()

	defer db.DBConn.Close()
	//start of the router
	r := mux.NewRouter()
	r.HandleFunc("/{key}", service.ProductHandler)

	r.HandleFunc("/product/{key}", service.CreateBlog).Methods("POST")
	r.HandleFunc("/product/{key}", service.ReadBlog).Methods("GET")
	r.HandleFunc("/product/{key}", service.UpdateBlog).Methods("PUT")
	r.HandleFunc("/product/{key}", service.DeleteBlog).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

	fmt.Println("listen to 8080")

	
}
