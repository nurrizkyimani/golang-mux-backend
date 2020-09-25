package service

import (
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"net/http"
	"github.com/nurrizkyimani/golang-mux-backend/model"
	"github.com/gorilla/mux"
	"fmt"
		"encoding/json"

	"io/ioutil"
)


//CreateBlog is a test
func CreateBlog(w http.ResponseWriter, r *http.Request) {
	db := db.DBConn

	var b model.Blog

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	res, _ := ioutil.ReadAll(r.Body)
	db.Create(&res)
}

//ReadBlog is cool
func ReadBlog(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)


	db := db.DBConn
	var blog model.Blog

	result := db.First(&blog, params["key"]) 
	json.NewEncoder(w).Encode(result)
}

//DeleteBlog  the blog
func DeleteBlog(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	db := db.DBConn

	db.Delete(&vars)
}	

//UpdateBlog the blog
func UpdateBlog(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	db := db.DBConn

	var b model.Blog

	db.First(&b, params["key"])

	_, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	b.Title = "new flex"
	b.Blog = "new desc"

	db.Save(&b)



}

//CreateTest is xx
func CreateTest(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	key := params["key"]

	db := db.DBConn
	blog := model.Blog {
		Title: key,
		Blog: "adfasdfadsf",
		Upvotes: 2 }
	db.Create(&blog)
}


// HomeHandler xxx
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["key"]

	fmt.Fprintf(w, "You've requested the book: %s on", key)
}


//ProductHandler is xxx
func ProductHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	key := vars["key"]

	for _, article := range model.Articles {
		if article.ID == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

//CommentCreated; 

//CreateComment xxx
func CreateComment() {

}

//ReadCommentById xxx
func ReadCommentById(){

}

//ReadCommentAll xxx
func ReadCommentAll(){

}

//UpdateCommentById xxx
func UpdateCommentById(){

}

//DeleteCommentById xxx
func DeleteCommentById(){

}
