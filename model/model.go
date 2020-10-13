package model

import (
	"github.com/jinzhu/gorm"
)

// Blog is a xxx
type Blog struct {
	gorm.Model
	Title       string
	Blog        string  
	Upvotes     int32
	BlogComment BlogComment `gorm:"embedded"`
}

//BlogComment is a xx
type BlogComment struct {
	gorm.Model
	Desc    string
	upvotes int32
	
}

//Article xxx
type Article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"content"`
}


type User struct {
	ID string `json:"ID"`
	Email string `json:"email"`
	Password string `json:"password"`
}



//Articles is xxx
var Articles =  []Article{
		Article{ID: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Article{ID: "3", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"}}

