package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"github.com/nurrizkyimani/golang-mux-backend/model"
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
func CreateComment(w http.ResponseWriter, r *http.Request) {

	db := db.DBConn

	_, err := ioutil.ReadAll(r.Body)


	if err != nil {
			log.Println("Error in : %v", err)
			http.Error(w, "can't read your request", http.StatusBadRequest)
			return
	}

	var comment model.BlogComment
	
	error := json.NewDecoder(r.Body).Decode(&comment)

	print("This si %v error ", error)

	print(comment)

	db.Create(&comment)

}

//ReadCommentByID xxx
func ReadCommentByID(w http.ResponseWriter, r *http.Request, ){
	w.Header().Set("Content-Type", "application/json")

	db := db.DBConn
	vars := mux.Vars(r)

	ID := vars["ID"]
	var comment model.BlogComment

	res := db.First(&comment, ID)
	json.NewEncoder(w).Encode(res)

}

//ReadCommentAll xxx
func ReadCommentAll(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	db := db.DBConn

	var comment model.BlogComment

	result := db.Find(&comment)


	json.NewDecoder(w).Decode(result)


}

//UpdateCommentByID xxx
func UpdateCommentByID(){
	
}

//DeleteCommentById xxx
func DeleteCommentById(){

}
