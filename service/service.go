package service

import (
	"github.com/nurrizkyimani/golang-mux-backend/db"
	"net/http"
	"github.com/nurrizkyimani/golang-mux-backend/model"
)


//CreateBlog is a test
func CreateBlog() {
	db := db.DBConn
	blog := Blog{}

	db.Create(&blog)
}

//ReadBlog is cool
func ReadBlog(w http.Response, r *http.Request){
	db := db.DBConn

	db.Get()
}
