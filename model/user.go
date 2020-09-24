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
	// BlogComment *BlogComment `gorm:"embedded`
}

//BlogComment is a xx
type BlogComment struct {
	gorm.Model
	Desc    string
	upvotes int32
	BlogID  uint
	Blog    Blog `gorm:"foreignKey:BlogID"`
}


