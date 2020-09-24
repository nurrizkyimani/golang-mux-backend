package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"github.com/nurrizkyimani/golang-mux-backend/service"

	// "gorm.io/gorm"s

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// this is comment to add
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintf(w, "You've requested the book: %s on", key)
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

type Article struct {
	ID      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}

var Articles []Article

func initDatabase() {
	var err error
	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	dbURL := os.Getenv("db_url")
	fmt.Println(dbURL)

	db.DBConn, err = gorm.Open("postgres", dbURL)
	if err != nil {
		fmt.Print(err)
		panic("failed to conenct ")
	}

	fmt.Println("Connection opened database")

	// database.DBConn.AutoMigrate(&film.Film{})/a

	fmt.Println("database migrated")

	// db = conn
	// db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

func main() {
	initDatabase()

	Articles = []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "3", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}

	r := mux.NewRouter()
	r.HandleFunc("/{key}", HomeHandler)
	r.HandleFunc("/product/{key}", ProductHandler)
	r.HandleFunc("/product/test/{key}", service.CreateTest)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))

	defer db.DBConn.Close()
}
