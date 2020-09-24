package service

import (
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"net/http"
	"github.com/nurrizkyimani/golang-mux-backend/model"
	"github.com/gorilla/mux"
)


//CreateBlog is a test
func CreateBlog() {
	db := db.DBConn
	blog := model.Blog {
		Title: "title 2",
		Blog: "adfasdfadsf",
		Upvotes: 2 }
	db.Create(&blog)
}

//ReadBlog is cool
func ReadBlog(w http.Response, r *http.Request){
	vars := mux.Vars(r)
	db := db.DBConn

	db.First(&vars)
}

//DeleteBlog  the blog
func DeleteBlog(w http.Response, r *http.Request){
	vars := mux.Vars(r)
	db := db.DBConn

	db.Delete(&vars)
}	

//UpdateBlog the blog
func UpdateBlog(w http.Response, r *http.Request){
	// vars := mux.Vars(r)

}

///CreateTest thisis 
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
